package usecase

import (
	"context"
	"time"

	_interface "main/features/sensors/model/interface"
	"main/features/sensors/model/request"
)

type SetThresholdSensorUseCase struct {
	Repository     _interface.ISetThresholdSensorRepository
	ContextTimeout time.Duration
}

func NewSetThresholdSensorUseCase(repository _interface.ISetThresholdSensorRepository, timeout time.Duration) _interface.ISetThresholdSensorUseCase {
	return &SetThresholdSensorUseCase{Repository: repository, ContextTimeout: timeout}
}

func (d *SetThresholdSensorUseCase) SetThresholdSensor(c context.Context, req *request.ReqSetThresholdSensor) error {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()

	thrsholdDTO := CreateThresholdDTO(req)
	err := d.Repository.UpdateOneThresholdSensor(ctx, &thrsholdDTO)
	if err != nil {
		return err
	}
	return nil
}
