package usecase

import (
	"context"
	"time"

	_interface "main/features/sensors/model/interface"
	"main/features/sensors/model/request"
	"main/features/sensors/model/response"
	"main/utils/db"
)

type GetSensorUseCase struct {
	Repository     _interface.IGetSensorRepository
	ContextTimeout time.Duration
}

func NewGetSensorUseCase(repository _interface.IGetSensorRepository, timeout time.Duration) _interface.IGetSensorUseCase {
	return &GetSensorUseCase{Repository: repository, ContextTimeout: timeout}
}

func (d *GetSensorUseCase) GetSensor(c context.Context, req *request.ReqGetSensor) (*response.ResGetSensor, error) {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()
	sensorDTO, err := d.Repository.FindOneSensor(ctx, req.SensorID)
	if err != nil {
		return nil, err
	}
	res := &response.ResGetSensor{
		SensorID:        sensorDTO.SensorID,
		MotionDetection: sensorDTO.MotionDetection,
		LightStatus:     sensorDTO.LightStatus,
	}
	sensorEventDTOList, err := d.Repository.FindAllSensorEvent(ctx, sensorDTO.SensorID)
	if err != nil {
		return nil, err
	}
	for _, sensorEventDTO := range sensorEventDTOList {
		if sensorEventDTO.Type == "light" {
			res.LightStatus = "error"
		}
		if sensorEventDTO.Type == "fire" {
			res.FireDetector = "detection"
		}
	}

	if sensorDTO.Position != (db.Position{}) {
		res.Position = response.Position{
			X: sensorDTO.Position.X,
			Y: sensorDTO.Position.Y,
		}
	}
	for _, sensor := range sensorDTO.Sensors {
		res.Sensors = append(res.Sensors, response.Sensor{
			Name:   sensor.Name,
			Value:  sensor.Value,
			Status: sensor.Status,
			Unit:   sensor.Unit,
		})
	}

	return res, nil
}
