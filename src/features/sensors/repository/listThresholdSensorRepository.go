package repository

import (
	"context"

	_interface "main/features/sensors/model/interface"
	"main/utils/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewListThresholdSensorRepository(mongoDB *mongo.Client) _interface.IListThresholdSensorRepository {
	return &ListThresholdSensorRepository{mongoDB: mongoDB}
}

func (r *ListThresholdSensorRepository) FindAllThresholdSensor(ctx context.Context) ([]db.SensorThresholdDTO, error) {
	cursor, err := db.SensorThresholdCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var thresholdList []db.SensorThresholdDTO
	if err := cursor.All(ctx, &thresholdList); err != nil {
		return nil, err
	}

	return thresholdList, nil
}
