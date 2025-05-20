package repository

import (
	_interface "main/features/sensors/model/interface"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewTopicRegisterSensorRepository(mongoDB *mongo.Client) _interface.ITopicRegisterSensorRepository {
	return &TopicRegisterSensorRepository{mongoDB: mongoDB}
}
