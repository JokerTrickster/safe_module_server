package repository

import (
	"context"

	_interface "main/features/sensors/model/interface"
	"main/utils/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewListSensorRepository(mongoDB *mongo.Client) _interface.IListSensorRepository {
	return &ListSensorRepository{mongoDB: mongoDB}
}

func (r *ListSensorRepository) FindAllSensor(ctx context.Context) ([]db.SensorDTO, error) {
	cursor, err := db.SensorsCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var sensorList []db.SensorDTO
	if err := cursor.All(ctx, &sensorList); err != nil {
		return nil, err
	}

	return sensorList, nil
}
