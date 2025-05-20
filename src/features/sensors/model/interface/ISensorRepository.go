package _interface

import (
	"context"
	"main/utils/db"
)

type IGetSensorRepository interface {
	FindOneSensor(ctx context.Context, macAddress string) (db.SensorDTO, error)
}

type ISetLightSensorRepository interface {
	UpdateOneLightSensor(ctx context.Context, sensorID string, status string) error
}

type ITopicRegisterSensorRepository interface {
}

type IGetLightSensorRepository interface {
	UpdateOneLightSensor(ctx context.Context, sensorID string, status string) error
}
