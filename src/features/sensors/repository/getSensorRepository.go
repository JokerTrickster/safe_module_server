package repository

import (
	"context"
	"fmt"

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
		fmt.Println("err", err)
		return db.SensorDTO{}, err
	}

	return sensor, nil
}

func (r *GetSensorRepository) FindAllSensorEvent(ctx context.Context, sensorID string) ([]db.SensorEventDTO, error) {
	filter := bson.M{"sensorID": sensorID, "confirmed": false}
	cursor, err := db.SensorEventsCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var sensorEventList []db.SensorEventDTO
	if err := cursor.All(ctx, &sensorEventList); err != nil {
		return nil, err
	}

	return sensorEventList, nil
}
