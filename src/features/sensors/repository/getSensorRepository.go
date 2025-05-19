package repository

import (
	"context"

	_interface "main/features/sensors/model/interface"
	"main/utils/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewGetSensorRepository(mongoDB *mongo.Client) _interface.IGetSensorRepository {
	return &GetSensorRepository{mongoDB: mongoDB}
}

func (r *GetSensorRepository) FindOneSensor(ctx context.Context, sensorID string) (db.SensorDTO, error) {

	filter := bson.M{"sensorID": sensorID}
	var sensor db.SensorDTO
	err := db.SensorsCollection.FindOne(ctx, filter).Decode(&sensor)
	if err != nil {
		return db.SensorDTO{}, err
	}

	return sensor, nil
}
