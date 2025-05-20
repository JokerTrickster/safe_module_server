package repository

import (
	"context"
	"time"

	_interface "main/features/sensors/model/interface"
	"main/utils/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewGetLightSensorRepository(mongoDB *mongo.Client) _interface.IGetLightSensorRepository {
	return &GetLightSensorRepository{mongoDB: mongoDB}
}

func (r *GetLightSensorRepository) UpdateOneLightSensor(ctx context.Context, sensorID string, status string) error {

	filter := bson.M{"sensorID": sensorID}
	// lightStatus , updatedAt 업데이트
	update := bson.M{"$set": bson.M{"lightStatus": status, "updatedAt": time.Now()}}

	_, err := db.SensorsCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil

}
