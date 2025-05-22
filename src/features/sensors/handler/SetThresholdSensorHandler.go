package handler

import (
	"context"
	"main/features/sensors/model/request"
	"main/utils"
	"net/http"

	_interface "main/features/sensors/model/interface"

	"github.com/labstack/echo/v4"
)

type SetThresholdSensorHandler struct {
	UseCase _interface.ISetThresholdSensorUseCase
}

func NewSetThresholdSensorHandler(c *echo.Echo, useCase _interface.ISetThresholdSensorUseCase) _interface.ISetThresholdSensorHandler {
	handler := &SetThresholdSensorHandler{
		UseCase: useCase,
	}

	c.POST("/v0.1/sensors/threshold", handler.SetThresholdSensor)

	return handler
}

// 센서 임계치 설정
// @Router /v0.1/sensors/threshold [post]
// @Summary 센서 임계치 설정
// @Description
// @Description ■ errCode with 400
// @Description PARAM_BAD : 파라미터 오류
// @Description USER_NOT_FOUND : 유저를 찾을 수 없음
// @Description
// @Description ■ errCode with 500
// @Description INTERNAL_SERVER : 내부 로직 처리 실패
// @Description INTERNAL_DB : DB 처리 실패
// @Param SetThresholdSensor body request.ReqSetThresholdSensor true "SetThresholdSensor"
// @Produce json
// @Success 200 {object} bool
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Tags sensors
func (h *SetThresholdSensorHandler) SetThresholdSensor(c echo.Context) error {
	req := request.ReqSetThresholdSensor{}
	if err := utils.ValidateReq(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, false)
	}

	err := h.UseCase.SetThresholdSensor(context.Background(), &req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, false)
	}

	return c.JSON(http.StatusOK, true)
}
