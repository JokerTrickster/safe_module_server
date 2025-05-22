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

type IListSensorRepository interface {
	FindAllSensor(ctx context.Context) ([]db.SensorDTO, error)
}

type ISetThresholdSensorRepository interface {
	UpdateOneThresholdSensor(ctx context.Context, thresholdDTO *db.SensorThresholdDTO) error
}
