package repository

import (
	"context"
	"fmt"
	_interface "main/features/sensors/model/interface"
	"main/utils/db"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewConfirmEventSensorRepository(mongoDB *mongo.Client) _interface.IConfirmEventSensorRepository {
	return &ConfirmEventSensorRepository{mongoDB: mongoDB}
}

func (r *ConfirmEventSensorRepository) UpdateOneConfirmEventSensor(ctx context.Context, sensorID, eventType, status string) error {
	// 3개의 필드가 모두 일치하는 문서 찾기
	filter := bson.M{
		"sensorID":  sensorID,
		"type":      eventType,
		"status":    status,
		"confirmed": false,
	}

	// confirmed를 true로 업데이트
	update := bson.M{
		"$set": bson.M{
			"confirmed": true,
			"updatedAt": time.Now(),
		},
	}

	// 업데이트 실행
	result, err := db.SensorEventsCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	// 업데이트된 문서가 없는 경우
	if result.MatchedCount == 0 {
		return fmt.Errorf("no document found with sensorID: %s, type: %s, status: %s", sensorID, eventType, status)
	}

	return nil
}
