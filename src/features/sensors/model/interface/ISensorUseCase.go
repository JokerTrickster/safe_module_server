package _interface

import (
	"context"
	"main/features/sensors/model/request"
	"main/features/sensors/model/response"
)

type IGetSensorUseCase interface {
	GetSensor(c context.Context, req *request.ReqGetSensor) (*response.ResGetSensor, error)
}
