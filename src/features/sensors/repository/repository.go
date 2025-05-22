package repository

import "go.mongodb.org/mongo-driver/mongo"

type GetSensorRepository struct {
	mongoDB *mongo.Client
}

type SetLightSensorRepository struct {
	mongoDB *mongo.Client
}

type TopicRegisterSensorRepository struct {
	mongoDB *mongo.Client
}

type GetLightSensorRepository struct {
	mongoDB *mongo.Client
}

type ListSensorRepository struct {
	mongoDB *mongo.Client
}

type SetThresholdSensorRepository struct {
	mongoDB *mongo.Client
}

type ListThresholdSensorRepository struct {
	mongoDB *mongo.Client
}

type SetPositionSensorRepository struct {
	mongoDB *mongo.Client
}
