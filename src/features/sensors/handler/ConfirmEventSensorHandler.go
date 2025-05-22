package handler

import (
	"context"
	"main/features/sensors/model/request"
	"main/utils"
	"net/http"

	_interface "main/features/sensors/model/interface"

	"github.com/labstack/echo/v4"
)

type ConfirmEventSensorHandler struct {
	UseCase _interface.IConfirmEventSensorUseCase
}

func NewConfirmEventSensorHandler(c *echo.Echo, useCase _interface.IConfirmEventSensorUseCase) _interface.IConfirmEventSensorHandler {
	handler := &ConfirmEventSensorHandler{
		UseCase: useCase,
	}

	c.PUT("/v0.1/sensors/event", handler.ConfirmEventSensor)

	return handler
}

// 이벤트 처리 확인
// @Router /v0.1/sensors/event [put]
// @Summary 이벤트 처리 확인
// @Description
// @Description ■ errCode with 400
// @Description PARAM_BAD : 파라미터 오류
// @Description USER_NOT_FOUND : 유저를 찾을 수 없음
// @Description
// @Description ■ errCode with 500
// @Description INTERNAL_SERVER : 내부 로직 처리 실패
// @Description INTERNAL_DB : DB 처리 실패
// @Param ConfirmEventSensor body request.ReqConfirmEventSensor true "ConfirmEventSensor"
// @Produce json
// @Success 200 {object} bool
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Tags sensors
func (h *ConfirmEventSensorHandler) ConfirmEventSensor(c echo.Context) error {
	req := request.ReqConfirmEventSensor{}
	if err := utils.ValidateReq(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, false)
	}

	err := h.UseCase.ConfirmEventSensor(context.Background(), &req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, false)
	}

	return c.JSON(http.StatusOK, true)
}
