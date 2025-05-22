package log

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"main/utils/db"
	"os"
	"path/filepath"
	"time"
)

const (
	LogBufferSize = 1000 // 버퍼 채널 크기
)

var (
	logChan chan map[string]interface{}
	dbLog   bool
	logFile *os.File
)

const (
	Info  = "INFO"
	Error = "ERROR"
	Debug = "DEBUG"
	Warn  = "WARN"
)

// InitLogger initializes the logger with buffer channel
func InitLogger(useDB bool) error {
	dbLog = useDB
	logChan = make(chan map[string]interface{}, LogBufferSize)

	if !dbLog {
		// 로그 파일 생성
		logDir := "logs"
		if err := os.MkdirAll(logDir, 0755); err != nil {
			return fmt.Errorf("failed to create log directory: %v", err)
		}

		logPath := filepath.Join(logDir, fmt.Sprintf("app_%s.log", time.Now().Format("2006-01-02")))
		var err error
		logFile, err = os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			return fmt.Errorf("failed to open log file: %v", err)
		}
	}

	go processLogs()
	fmt.Println("로그 초기화 완료!")
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

// processLogs processes logs from the buffer channel and saves to MongoDB or file
func processLogs() {
	batch := make([]interface{}, 0, 100) // 배치 크기 100
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case logEntry := <-logChan:
			if dbLog {
				batch = append(batch, logEntry)
				if len(batch) >= 100 {
					saveLogsToMongoDB(batch)
					batch = batch[:0]
				}
			} else {
				saveLogToFile(logEntry)
			}

		case <-ticker.C:
			if dbLog && len(batch) > 0 {
				saveLogsToMongoDB(batch)
				batch = batch[:0]
			}
		}
	}
}

// saveLogToFile saves a single log entry to file
func saveLogToFile(logEntry map[string]interface{}) {
	if logFile == nil {
		return
	}

	logEntry["timestamp"] = time.Now().Format(time.RFC3339)
	jsonData, err := json.Marshal(logEntry)
	if err != nil {
		log.Printf("Error marshaling log entry: %v", err)
		return
	}

	if _, err := logFile.Write(append(jsonData, '\n')); err != nil {
		log.Printf("Error writing to log file: %v", err)
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

	if logFile != nil {
		logFile.Close()
	}
}
