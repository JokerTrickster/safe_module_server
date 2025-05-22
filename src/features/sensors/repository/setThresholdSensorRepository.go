package repository

import (
	"context"
	_interface "main/features/sensors/model/interface"
	"main/utils/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewSetThresholdSensorRepository(mongoDB *mongo.Client) _interface.ISetThresholdSensorRepository {
	return &SetThresholdSensorRepository{mongoDB: mongoDB}
}

func (r *SetThresholdSensorRepository) UpdateOneThresholdSensor(ctx context.Context, thresholdDTO *db.SensorThresholdDTO) error {
	filter := bson.M{"name": thresholdDTO.Name}

	// 기존 문서 확인
	var existingDoc db.SensorThresholdDTO
	err := db.SensorThresholdCollection.FindOne(ctx, filter).Decode(&existingDoc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// 문서가 없으면 새로운 문서 삽입
			_, err = db.SensorThresholdCollection.InsertOne(ctx, thresholdDTO)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}

	// 문서가 있으면 threshold 값만 업데이트
	update := bson.M{"$set": bson.M{"threshold": thresholdDTO.Threshold}}
	_, err = db.SensorThresholdCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}
