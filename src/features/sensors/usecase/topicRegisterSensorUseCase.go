package usecase

import (
	"context"
	"time"

	_interface "main/features/sensors/model/interface"
	"main/features/sensors/model/request"
	"main/utils/mqtt"
)

type TopicRegisterSensorUseCase struct {
	Repository     _interface.ITopicRegisterSensorRepository
	ContextTimeout time.Duration
}

func NewTopicRegisterSensorUseCase(repository _interface.ITopicRegisterSensorRepository, timeout time.Duration) _interface.ITopicRegisterSensorUseCase {
	return &TopicRegisterSensorUseCase{Repository: repository, ContextTimeout: timeout}
}

func (d *TopicRegisterSensorUseCase) TopicRegisterSensor(c context.Context, req *request.ReqTopicRegisterSensor) error {
	_, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()

	err := mqtt.Subscribe(req.Topic, byte(req.Qos), mqtt.SensorLightResponseHandler)
	if err != nil {
		return err
	}

	return nil
}
