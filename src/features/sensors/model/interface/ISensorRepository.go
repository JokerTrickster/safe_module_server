package _interface

import (
	"context"
	"main/utils/db"
)

type IGetSensorRepository interface {
	FindOneSensor(ctx context.Context, macAddress string) (db.SensorDTO, error)
}
