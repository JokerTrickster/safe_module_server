package handler

import (
	"context"
	"main/features/sensors/model/request"
	"main/utils"
	"net/http"

	_interface "main/features/sensors/model/interface"

	"github.com/labstack/echo/v4"
)

type SetPositionSensorHandler struct {
	UseCase _interface.ISetPositionSensorUseCase
}

func NewSetPositionSensorHandler(c *echo.Echo, useCase _interface.ISetPositionSensorUseCase) _interface.ISetPositionSensorHandler {
	handler := &SetPositionSensorHandler{
		UseCase: useCase,
	}

	c.PUT("/v0.1/sensors", handler.SetPositionSensor)

	return handler
}

// 센서 위치 설정
// @Router /v0.1/sensors [put]
// @Summary 센서 위치 설정
// @Description
// @Description ■ errCode with 400
// @Description PARAM_BAD : 파라미터 오류
// @Description USER_NOT_FOUND : 유저를 찾을 수 없음
// @Description
// @Description ■ errCode with 500
// @Description INTERNAL_SERVER : 내부 로직 처리 실패
// @Description INTERNAL_DB : DB 처리 실패
// @Param SetPositionSensor body request.ReqSetPositionSensor true "SetPositionSensor"
// @Produce json
// @Success 200 {object} bool
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Tags sensors
func (h *SetPositionSensorHandler) SetPositionSensor(c echo.Context) error {
	req := request.ReqSetPositionSensor{}
	if err := utils.ValidateReq(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, false)
	}

	err := h.UseCase.SetPositionSensor(context.Background(), &req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, false)
	}

	return c.JSON(http.StatusOK, true)
}
