package handler

import (
	"context"
	"main/features/sensors/model/request"
	"main/utils"
	"net/http"

	_interface "main/features/sensors/model/interface"

	"github.com/labstack/echo/v4"
)

type GetLightSensorHandler struct {
	UseCase _interface.IGetLightSensorUseCase
}

func NewGetLightSensorHandler(c *echo.Echo, useCase _interface.IGetLightSensorUseCase) _interface.IGetLightSensorHandler {
	handler := &GetLightSensorHandler{
		UseCase: useCase,
	}

	c.GET("/v0.1/light/status", handler.GetLightSensor)

	return handler
}

// 조명 정보 가져오기
// @Router /v0.1/light/status [get]
// @Summary 조명 정보 가져오기
// @Description
// @Description ■ errCode with 400
// @Description PARAM_BAD : 파라미터 오류
// @Description USER_NOT_FOUND : 유저를 찾을 수 없음
// @Description
// @Description ■ errCode with 500
// @Description INTERNAL_SERVER : 내부 로직 처리 실패
// @Description INTERNAL_DB : DB 처리 실패
// @Param sensorID query string true "sensorID"
// @Produce json
// @Success 200 {object} response.ResGetLightSensor
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Tags sensors
func (h *GetLightSensorHandler) GetLightSensor(c echo.Context) error {
	req := request.ReqGetLightSensor{}
	if err := utils.ValidateReq(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	res, err := h.UseCase.GetLightSensor(context.Background(), &req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, res)
}
