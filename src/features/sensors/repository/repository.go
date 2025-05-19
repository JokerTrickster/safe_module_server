package repository

import "go.mongodb.org/mongo-driver/mongo"

type GetSensorRepository struct {
	mongoDB *mongo.Client
}

type SetLightSensorRepository struct {
	mongoDB *mongo.Client
}
