package handler

import (
	"context"
	"main/features/sensors/model/request"
	"main/features/sensors/model/response"
	"main/utils"
	"net/http"

	_interface "main/features/sensors/model/interface"

	"github.com/labstack/echo/v4"
)

type SetLightSensorHandler struct {
	UseCase _interface.ISetLightSensorUseCase
}

func NewSetLightSensorHandler(c *echo.Echo, useCase _interface.ISetLightSensorUseCase) _interface.ISetLightSensorHandler {
	handler := &SetLightSensorHandler{
		UseCase: useCase,
	}

	c.POST("/v0.1/sensors/light", handler.SetLightSensor)

	return handler
}

// 센서 조명 켜기/끄기
// @Router /v0.1/sensors/light [post]
// @Summary 센서 조명 켜기/끄기
// @Description
// @Description ■ errCode with 400
// @Description PARAM_BAD : 파라미터 오류
// @Description USER_NOT_FOUND : 유저를 찾을 수 없음
// @Description
// @Description ■ errCode with 500
// @Description INTERNAL_SERVER : 내부 로직 처리 실패
// @Description INTERNAL_DB : DB 처리 실패
// @Param setLightSensor body request.ReqSetLightSensor true "setLightSensor"
// @Produce json
// @Success 200 {object} response.ResSetLightSensor
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Tags sensors
func (h *SetLightSensorHandler) SetLightSensor(c echo.Context) error {
	req := request.ReqSetLightSensor{}
	if err := utils.ValidateReq(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, false)
	}

	err := h.UseCase.SetLightSensor(context.Background(), &req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, false)
	}

	return c.JSON(http.StatusOK, response.ResSetLightSensor{
		SensorID:    req.SensorID,
		LightStatus: req.Status,
	})
}
