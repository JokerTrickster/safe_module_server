package _interface

import (
	"context"
	"main/features/sensors/model/request"
	"main/features/sensors/model/response"
)

type IGetSensorUseCase interface {
	GetSensor(c context.Context, req *request.ReqGetSensor) (*response.ResGetSensor, error)
}

type ISetLightSensorUseCase interface {
	SetLightSensor(c context.Context, req *request.ReqSetLightSensor) error
}

type ITopicRegisterSensorUseCase interface {
	TopicRegisterSensor(c context.Context, req *request.ReqTopicRegisterSensor) error
}

type IGetLightSensorUseCase interface {
	GetLightSensor(c context.Context, req *request.ReqGetLightSensor) (*response.ResGetLightSensor, error)
}

type IListSensorUseCase interface {
	ListSensor(c context.Context) (*response.ResListSensor, error)
}

type ISetThresholdSensorUseCase interface {
	SetThresholdSensor(c context.Context, req *request.ReqSetThresholdSensor) error
}

type IListThresholdSensorUseCase interface {
	ListThresholdSensor(c context.Context) (*response.ResListThresholdSensor, error)
}
