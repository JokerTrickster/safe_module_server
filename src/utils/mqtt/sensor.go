package mqtt

import (
	"context"
	"encoding/json"
	"main/utils/db"
	_log "main/utils/log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SensorData struct {
	SensorID    string `json:"sensor_id"`
	BootingTime int64  `json:"booting_time"`
	SensorList  []struct {
		Name    string  `json:"name"`
		Status  string  `json:"status"`
		Value   float64 `json:"value"`
		Unit    string  `json:"unit"`
		RawData []int   `json:"raw_data"`
	} `json:"sensor_list"`
}

// 센서 데이터 핸들러
func SensorDataHandler(client mqtt.Client, msg mqtt.Message) {
	_log.Log(_log.Info, "Received message on topic!!", map[string]interface{}{
		"topic":   msg.Topic(),
		"payload": string(msg.Payload()),
	})

	// JSON 파싱
	var sensorData SensorData
	if err := json.Unmarshal(msg.Payload(), &sensorData); err != nil {
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

	// 센서 데이터 변환
	for i, sensor := range sensorData.SensorList {
		sensorDTO.Sensors[i] = db.Sensor{
			Name:   sensor.Name,
			Value:  sensor.Value,
			Status: sensor.Status,
		}
	}

	// MongoDB에 저장 (upsert)
	ctx := context.Background()
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
		sensorDTO.LightStatus = existingDoc.LightStatus
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

	_log.Log(_log.Info, "Successfully saved sensor data", map[string]interface{}{
		"sensorID": sensorDTO.SensorID,
	})
}
