package usecase

import (
	"context"
	"time"

	_interface "main/features/sensors/model/interface"
	"main/features/sensors/model/request"
)

type SetPositionSensorUseCase struct {
	Repository     _interface.ISetPositionSensorRepository
	ContextTimeout time.Duration
}

func NewSetPositionSensorUseCase(repository _interface.ISetPositionSensorRepository, timeout time.Duration) _interface.ISetPositionSensorUseCase {
	return &SetPositionSensorUseCase{Repository: repository, ContextTimeout: timeout}
}

func (d *SetPositionSensorUseCase) SetPositionSensor(c context.Context, req *request.ReqSetPositionSensor) error {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()

	err := d.Repository.UpdateOnePositionSensor(ctx, req.SensorID, req.Position)
	if err != nil {
		return err
	}

	return nil
}
