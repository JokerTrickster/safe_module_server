package repository

import (
	"context"
	_interface "main/features/sensors/model/interface"
	"main/features/sensors/model/request"
	"main/utils/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewSetPositionSensorRepository(mongoDB *mongo.Client) _interface.ISetPositionSensorRepository {
	return &SetPositionSensorRepository{mongoDB: mongoDB}
}

func (r *SetPositionSensorRepository) UpdateOnePositionSensor(ctx context.Context, sensorID string, position request.Position) error {
	filter := bson.M{"sensorID": sensorID}
	update := bson.M{"$set": bson.M{"position": position}}
	_, err := db.SensorsCollection.UpdateOne(ctx, filter, update)
	return err
}
