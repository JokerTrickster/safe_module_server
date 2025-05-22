package usecase

import (
	"context"
	"time"

	_interface "main/features/sensors/model/interface"
	"main/features/sensors/model/request"
)

type ConfirmEventSensorUseCase struct {
	Repository     _interface.IConfirmEventSensorRepository
	ContextTimeout time.Duration
}

func NewConfirmEventSensorUseCase(repository _interface.IConfirmEventSensorRepository, timeout time.Duration) _interface.IConfirmEventSensorUseCase {
	return &ConfirmEventSensorUseCase{Repository: repository, ContextTimeout: timeout}
}

func (d *ConfirmEventSensorUseCase) ConfirmEventSensor(c context.Context, req *request.ReqConfirmEventSensor) error {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()

	err := d.Repository.UpdateOneConfirmEventSensor(ctx, req.SensorID, req.Type, req.Status)
	if err != nil {
		return err
	}

	return nil
}
