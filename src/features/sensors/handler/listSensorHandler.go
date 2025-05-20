package handler

import (
	"context"
	"net/http"

	_interface "main/features/sensors/model/interface"

	"github.com/labstack/echo/v4"
)

type ListSensorHandler struct {
	UseCase _interface.IListSensorUseCase
}

func NewListSensorHandler(c *echo.Echo, useCase _interface.IListSensorUseCase) _interface.IListSensorHandler {
	handler := &ListSensorHandler{
		UseCase: useCase,
	}

	c.GET("/v0.1/sensors/list", handler.ListSensor)

	return handler
}

// 센서 리스트 정보 가져오기
// @Router /v0.1/sensors/list [get]
// @Summary 센서 리스트 정보 가져오기
// @Description
// @Description ■ errCode with 400
// @Description PARAM_BAD : 파라미터 오류
// @Description USER_NOT_FOUND : 유저를 찾을 수 없음
// @Description
// @Description ■ errCode with 500
// @Description INTERNAL_SERVER : 내부 로직 처리 실패
// @Description INTERNAL_DB : DB 처리 실패
// @Produce json
// @Success 200 {object} response.ResListSensor
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Tags sensors
func (h *ListSensorHandler) ListSensor(c echo.Context) error {

	res, err := h.UseCase.ListSensor(context.Background())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}
