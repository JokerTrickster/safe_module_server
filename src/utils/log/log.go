package log

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"main/utils/db"
	"time"
)

const (
	LogBufferSize = 1000 // 버퍼 채널 크기
)

var (
	logChan chan map[string]interface{}
)

const (
	Info  = "INFO"
	Error = "ERROR"
	Debug = "DEBUG"
	Warn  = "WARN"
)

// InitLogger initializes the logger with buffer channel
func InitLogger() error {
	logChan = make(chan map[string]interface{}, LogBufferSize)
	go processLogs()

	fmt.Println("로그 초기화 완료 !")
	return nil
}

// Log writes a log entry to the buffer channel
func Log(level string, message string, data map[string]interface{}) {
	logEntry := map[string]interface{}{
		"timestamp": time.Now(),
		"level":     level,
		"message":   message,
	}

	// 추가 데이터가 있다면 병합
	if data != nil {
		for k, v := range data {
			logEntry[k] = v
		}
	}

	// 채널이 가득 찼을 경우를 대비한 non-blocking send
	select {
	case logChan <- logEntry:
		// 로그가 성공적으로 버퍼에 추가됨
	default:
		// 버퍼가 가득 찼을 경우 콘솔에 경고 출력
		log.Printf("Warning: Log buffer is full, dropping log entry: %s", message)
	}
}

// processLogs processes logs from the buffer channel and saves to MongoDB
func processLogs() {
	batch := make([]interface{}, 0, 100) // 배치 크기 100
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case logEntry := <-logChan:
			batch = append(batch, logEntry)

			// 배치가 가득 찼거나 채널이 비어있을 때 MongoDB에 저장
			if len(batch) >= 100 {
				saveLogsToMongoDB(batch)
				batch = batch[:0]
			}

		case <-ticker.C:
			// 주기적으로 남은 로그 저장
			if len(batch) > 0 {
				saveLogsToMongoDB(batch)
				batch = batch[:0]
			}
		}
	}
}

// saveLogsToMongoDB saves a batch of logs to MongoDB
func saveLogsToMongoDB(logs []interface{}) {
	if len(logs) == 0 {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// MongoDB에 배치 삽입
	_, err := db.LogsCollection.InsertMany(ctx, logs)
	if err != nil {
		log.Printf("Error saving logs to MongoDB: %v", err)
		// 실패한 로그를 JSON으로 출력
		for _, logEntry := range logs {
			if jsonData, err := json.Marshal(logEntry); err == nil {
				log.Printf("Failed log entry: %s", string(jsonData))
			}
		}
	}
}

// CloseLogger closes the logger and processes remaining logs
func CloseLogger() {
	// 남은 로그 처리
	if len(logChan) > 0 {
		log.Printf("Processing %d remaining logs...", len(logChan))
		time.Sleep(2 * time.Second) // 남은 로그 처리 대기
	}
	close(logChan)
}
