package usecase

import (
	"context"
	"time"

	_interface "main/features/sensors/model/interface"
	"main/features/sensors/model/response"
)

type ListThresholdSensorUseCase struct {
	Repository     _interface.IListThresholdSensorRepository
	ContextTimeout time.Duration
}

func NewListThresholdSensorUseCase(repository _interface.IListThresholdSensorRepository, timeout time.Duration) _interface.IListThresholdSensorUseCase {
	return &ListThresholdSensorUseCase{Repository: repository, ContextTimeout: timeout}
}

func (d *ListThresholdSensorUseCase) ListThresholdSensor(c context.Context) (*response.ResListThresholdSensor, error) {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()

	thresholdDTOList, err := d.Repository.FindAllThresholdSensor(ctx)
	if err != nil {
		return nil, err
	}
	res := &response.ResListThresholdSensor{
		ThresholdList: make([]response.Threshold, 0),
	}

	for _, thresholdDTO := range thresholdDTOList {
		res.ThresholdList = append(res.ThresholdList, response.Threshold{
			Name:      thresholdDTO.Name,
			Unit:      thresholdDTO.Unit,
			Threshold: int(thresholdDTO.Threshold),
		})
	}

	return res, nil
}
