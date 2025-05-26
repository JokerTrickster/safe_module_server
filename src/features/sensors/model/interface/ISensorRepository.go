package _interface

import (
	"context"
	"main/features/sensors/model/request"
	"main/utils/db"
)

type IGetSensorRepository interface {
	FindOneSensor(ctx context.Context, macAddress string) (db.SensorDTO, error)
	FindAllSensorEvent(ctx context.Context, sensorID string) ([]db.SensorEventDTO, error)
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
	FindAllSensorEvent(ctx context.Context, sensorID string) ([]db.SensorEventDTO, error)
}

type ISetThresholdSensorRepository interface {
	UpdateOneThresholdSensor(ctx context.Context, thresholdDTO *db.SensorThresholdDTO) error
}

type IListThresholdSensorRepository interface {
	FindAllThresholdSensor(ctx context.Context) ([]db.SensorThresholdDTO, error)
}

type ISetPositionSensorRepository interface {
	UpdateOnePositionSensor(ctx context.Context, sensorID string, position request.Position) error
}

type IConfirmEventSensorRepository interface {
	UpdateOneConfirmEventSensor(ctx context.Context, sensorID, eventType, status string) error
}
