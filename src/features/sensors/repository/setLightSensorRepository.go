package repository

import (
	"context"
	"fmt"
	_interface "main/features/sensors/model/interface"
	"main/utils/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewSetLightSensorRepository(mongoDB *mongo.Client) _interface.ISetLightSensorRepository {
	return &SetLightSensorRepository{mongoDB: mongoDB}
}

func (r *SetLightSensorRepository) UpdateOneLightSensor(ctx context.Context, sensorID string, status bool) error {
	//status pointer로 변경
	fmt.Println("UpdateOneLightSensor", sensorID, status)
	filter := bson.M{"sensorID": sensorID}
	statusPtr := &status
	update := bson.M{"$set": bson.M{"lightStatus": statusPtr}}
	_, err := db.SensorsCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}
