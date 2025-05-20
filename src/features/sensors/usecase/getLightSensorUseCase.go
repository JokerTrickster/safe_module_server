package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	_interface "main/features/sensors/model/interface"
	"main/features/sensors/model/request"
	"main/features/sensors/model/response"
	"main/utils/mqtt"

	"github.com/google/uuid"
)

type GetLightSensorUseCase struct {
	Repository     _interface.IGetLightSensorRepository
	ContextTimeout time.Duration
}

func NewGetLightSensorUseCase(repository _interface.IGetLightSensorRepository, timeout time.Duration) _interface.IGetLightSensorUseCase {
	return &GetLightSensorUseCase{Repository: repository, ContextTimeout: timeout}
}

func (d *GetLightSensorUseCase) GetLightSensor(c context.Context, req *request.ReqGetLightSensor) (*response.ResGetLightSensor, error) {
	ctx, cancel := context.WithTimeout(c, d.ContextTimeout)
	defer cancel()

	// uuid 생성
	uuid := uuid.New().String()
	requestTopic := fmt.Sprintf("/control/light/request/get/%s", req.SensorID)
	responseTopic := fmt.Sprintf("/control/light/response/get/%s", req.SensorID)

	// json 형식을 만들거다. true 일 경우 status on, false 일 경우 status off
	var jsonData []byte
	jsonData, _ = json.Marshal(map[string]interface{}{})

	resp, err := mqtt.PublishAndWaitForResponse(requestTopic, 2, jsonData, uuid, responseTopic, 5*time.Second)
	if err != nil {
		return nil, err
	}

	// 응답 페이로드 파싱
	var lightResp response.LightResponse
	if err := json.Unmarshal(resp.Payload, &lightResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	err = d.Repository.UpdateOneLightSensor(ctx, req.SensorID, lightResp.Status)
	if err != nil {
		return nil, err
	}

	res := &response.ResGetLightSensor{
		Status: lightResp.Status,
	}

	return res, nil
}
