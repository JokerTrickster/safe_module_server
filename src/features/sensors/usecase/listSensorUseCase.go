package usecase

import (
	"context"
	"fmt"
	"time"

	_interface "main/features/sensors/model/interface"
	"main/features/sensors/model/response"
)

type ListSensorUseCase struct {
	Repository     _interface.IListSensorRepository
	ContextTimeout time.Duration
}

func NewListSensorUseCase(repository _interface.IListSensorRepository, timeout time.Duration) _interface.IListSensorUseCase {
	return &ListSensorUseCase{Repository: repository, ContextTimeout: timeout}
}

func (d *ListSensorUseCase) ListSensor(c context.Context) (*response.ResListSensor, error) {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()

	sensorDTOList, err := d.Repository.FindAllSensor(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println(sensorDTOList)
	res := &response.ResListSensor{
		SensorList: make([]response.ListSensor, 0),
	}

	for i, sensorDTO := range sensorDTOList {
		res.SensorList = append(res.SensorList, response.ListSensor{
			SensorID:     sensorDTO.SensorID,
			LightStatus:  sensorDTO.LightStatus,
			FireDetector: sensorDTO.FireDetector,
		})
		resSensor := make([]response.Sensor, 0)
		for _, sensor := range sensorDTO.Sensors {
			tmpSensor := response.Sensor{
				Name:   sensor.Name,
				Value:  sensor.Value,
				Status: sensor.Status,
				Unit:   sensor.Unit,
			}
			resSensor = append(resSensor, tmpSensor)
		}
		res.SensorList[i].Sensors = resSensor
	}

	return res, nil
}
