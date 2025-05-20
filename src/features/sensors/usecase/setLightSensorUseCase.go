package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	_interface "main/features/sensors/model/interface"
	"main/features/sensors/model/request"
	"main/utils/mqtt"

	"github.com/google/uuid"
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

	// uuid 생성
	uuid := uuid.New().String()
	requestTopic := fmt.Sprintf("/control/light/request/set/%s", req.SensorID)
	responseTopic := fmt.Sprintf("/control/light/response/set/%s", req.SensorID)

	// json 형식을 만들거다. true 일 경우 status on, false 일 경우 status off
	var jsonData []byte
	jsonData, _ = json.Marshal(map[string]interface{}{
		"status": req.Status,
	})

	resp, err := mqtt.PublishAndWaitForResponse(requestTopic, 2, jsonData, uuid, responseTopic, 5*time.Second)
	if err != nil {
		return err
	}

	fmt.Println("응답 메시지 , ", string(resp.Payload))

	// db 상태 변경
	err = d.Repository.UpdateOneLightSensor(ctx, req.SensorID, req.Status)
	if err != nil {
		return err
	}

	return nil
}
