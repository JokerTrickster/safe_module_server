package mqtt

import (
	"context"
	"encoding/json"
	"main/utils/db"
	"sync"
	"time"

	_log "main/utils/log"

	"github.com/eclipse/paho.golang/paho"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SensorData struct {
	SensorID     string `json:"sensor_id"`
	BootingTime  int64  `json:"booting_time"`
	FireDetector string `json:"fire_detector"`
	LightStatus  string `json:"light_status"`
	SensorList   []struct {
		Name    string  `json:"name"`
		Status  string  `json:"status"`
		Value   float64 `json:"value"`
		Unit    string  `json:"unit"`
		RawData []int   `json:"raw_data"`
	} `json:"sensor_list"`
}

// 구독 상태를 관리할 sync.Map으로 변경
var subscribedTopics sync.Map

// 센서 데이터 핸들러
func SensorDataHandler(p *paho.Publish) {
	_log.Log(_log.Info, "Received message on topic!!", map[string]interface{}{
		"payload": string(p.Payload),
	})

	// JSON 파싱
	var sensorData SensorData
	if err := json.Unmarshal(p.Payload, &sensorData); err != nil {
		_log.Log(_log.Error, "Failed to parse sensor data", map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	createdAt := time.Now()
	updatedAt := time.Now()

	// SensorDTO로 변환
	sensorDTO := db.SensorDTO{
		SensorID:  sensorData.SensorID,
		Sensors:   make([]db.Sensor, len(sensorData.SensorList)),
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
	}

	ctx := context.Background()

	if sensorData.LightStatus == "shutdown" {
		DangerLightEventUseCase(ctx, sensorDTO)
	}
	if sensorData.FireDetector == "detection" {
		DangerFireEventUseCase(ctx, sensorDTO)
	}

	_log.Log(_log.Info, "Successfully saved sensor data", map[string]interface{}{
		"sensorID": sensorDTO.SensorID,
	})

	// 센서 데이터 변환
	for i, sensor := range sensorData.SensorList {
		sensorDTO.Sensors[i] = db.Sensor{
			Name:   sensor.Name,
			Value:  sensor.Value,
			Status: sensor.Status,
			Unit:   sensor.Unit,
		}
	}

	// MongoDB에 저장 (upsert)
	filter := bson.M{"sensorID": sensorDTO.SensorID}

	// 기존 문서 확인
	var existingDoc db.SensorDTO
	err := db.SensorsCollection.FindOne(ctx, filter).Decode(&existingDoc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// 문서가 없으면 createdAt 설정
			createdAt := time.Now()
			sensorDTO.CreatedAt = &createdAt
		} else {
			_log.Log(_log.Error, "Failed to check existing document", map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
	} else {
		// 기존 문서가 있으면 createdAt 유지
		sensorDTO.CreatedAt = existingDoc.CreatedAt
		sensorDTO.Position = existingDoc.Position
	}

	// updatedAt은 항상 현재 시간으로 설정
	sensorDTO.UpdatedAt = &updatedAt

	update := bson.M{"$set": sensorDTO}
	opts := options.Update().SetUpsert(true)

	_, err = db.SensorsCollection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		_log.Log(_log.Error, "Failed to save sensor data to MongoDB", map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

}

func DangerLightEventUseCase(ctx context.Context, sensorDTO db.SensorDTO) {
	now := time.Now()
	SensorEventDTO := db.SensorEventDTO{
		Type:      "light",
		Status:    "shutdown",
		SensorID:  sensorDTO.SensorID,
		Confirmed: false,
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	// 기존 문서 확인
	filter := bson.M{
		"type":      SensorEventDTO.Type,
		"status":    SensorEventDTO.Status,
		"sensorID":  SensorEventDTO.SensorID,
		"confirmed": SensorEventDTO.Confirmed,
	}

	// 문서 존재 여부 확인
	var existingDoc db.SensorEventDTO
	err := db.SensorEventsCollection.FindOne(ctx, filter).Decode(&existingDoc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// 문서가 없으면 새로 생성
			_, err = db.SensorEventsCollection.InsertOne(ctx, SensorEventDTO)
			if err != nil {
				_log.Log(_log.Error, "Failed to create sensor event", map[string]interface{}{
					"error": err.Error(),
				})
			}
		} else {
			_log.Log(_log.Error, "Failed to check existing sensor event", map[string]interface{}{
				"error": err.Error(),
			})
		}
	}
}

func DangerFireEventUseCase(ctx context.Context, sensorDTO db.SensorDTO) {
	now := time.Now()
	SensorEventDTO := db.SensorEventDTO{
		Type:      "fire",
		Status:    "detection",
		SensorID:  sensorDTO.SensorID,
		Confirmed: false,
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	// 기존 문서 확인
	filter := bson.M{
		"type":      SensorEventDTO.Type,
		"status":    SensorEventDTO.Status,
		"sensorID":  SensorEventDTO.SensorID,
		"confirmed": SensorEventDTO.Confirmed,
	}

	// 문서 존재 여부 확인
	var existingDoc db.SensorEventDTO
	err := db.SensorEventsCollection.FindOne(ctx, filter).Decode(&existingDoc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// 문서가 없으면 새로 생성
			_, err = db.SensorEventsCollection.InsertOne(ctx, SensorEventDTO)
			if err != nil {
				_log.Log(_log.Error, "Failed to create sensor event", map[string]interface{}{
					"error": err.Error(),
				})
			}
		} else {
			_log.Log(_log.Error, "Failed to check existing sensor event", map[string]interface{}{
				"error": err.Error(),
			})
		}
	}
}
