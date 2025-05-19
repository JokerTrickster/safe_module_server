package usecase

import (
	"context"
	"time"

	_interface "main/features/sensors/model/interface"
	"main/features/sensors/model/request"
)

type SetLightSensorUseCase struct {
	Repository     _interface.ISetLightSensorRepository
	ContextTimeout time.Duration
}

func NewSetLightSensorUseCase(repository _interface.ISetLightSensorRepository, timeout time.Duration) _interface.ISetLightSensorUseCase {
	return &SetLightSensorUseCase{Repository: repository, ContextTimeout: timeout}
}

func (d *SetLightSensorUseCase) SetLightSensor(c context.Context, req *request.ReqSetLightSensor) error {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()

	// db 상태 변경
	err := d.Repository.UpdateOneLightSensor(ctx, req.SensorID, req.Status)
	if err != nil {
		return err
	}

	return nil
}
