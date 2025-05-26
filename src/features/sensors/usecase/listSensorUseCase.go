package usecase

import (
	"context"
	"time"

	_interface "main/features/sensors/model/interface"
	"main/features/sensors/model/response"
	"main/utils/db"
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
	res := &response.ResListSensor{
		SensorList: make([]response.ListSensor, 0),
	}

	for _, sensorDTO := range sensorDTOList {
		tmpListSensor := response.ListSensor{
			SensorID: sensorDTO.SensorID,
		}
		sensorEventDTOList, err := d.Repository.FindAllSensorEvent(ctx, sensorDTO.SensorID)
		if err != nil {
			return nil, err
		}
		for _, sensorEventDTO := range sensorEventDTOList {
			if sensorEventDTO.Type == "light" {
				tmpListSensor.LightStatus = "shutdown"
			}
			if sensorEventDTO.Type == "fire" {
				tmpListSensor.FireDetector = "detection"
			}
		}

		if sensorDTO.Position != (db.Position{}) {
			tmpPosition := response.Position{
				X: sensorDTO.Position.X,
				Y: sensorDTO.Position.Y,
			}
			tmpListSensor.Position = tmpPosition
		}
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
		tmpListSensor.Sensors = resSensor
		res.SensorList = append(res.SensorList, tmpListSensor)
	}

	return res, nil
}
