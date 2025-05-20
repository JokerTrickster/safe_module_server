package handler

import (
	"context"
	"main/features/sensors/model/request"
	"main/utils"
	"net/http"

	_interface "main/features/sensors/model/interface"

	"github.com/labstack/echo/v4"
)

type TopicRegisterSensorHandler struct {
	UseCase _interface.ITopicRegisterSensorUseCase
}

func NewTopicRegisterSensorHandler(c *echo.Echo, useCase _interface.ITopicRegisterSensorUseCase) _interface.ITopicRegisterSensorHandler {
	handler := &TopicRegisterSensorHandler{
		UseCase: useCase,
	}

	c.POST("/v0.1/topic/register", handler.TopicRegisterSensor)

	return handler
}

// 토픽 구독 등록
// @Router /v0.1/topic/register [post]
// @Summary 토픽 구독 등록
// @Description
// @Description ■ errCode with 400
// @Description PARAM_BAD : 파라미터 오류
// @Description USER_NOT_FOUND : 유저를 찾을 수 없음
// @Description
// @Description ■ errCode with 500
// @Description INTERNAL_SERVER : 내부 로직 처리 실패
// @Description INTERNAL_DB : DB 처리 실패
// @Param TopicRegisterSensor body request.ReqTopicRegisterSensor true "TopicRegisterSensor"
// @Produce json
// @Success 200 {object} bool
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Tags sensors
func (h *TopicRegisterSensorHandler) TopicRegisterSensor(c echo.Context) error {
	req := request.ReqTopicRegisterSensor{}
	if err := utils.ValidateReq(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, false)
	}

	err := h.UseCase.TopicRegisterSensor(context.Background(), &req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, false)
	}

	return c.JSON(http.StatusOK, true)
}
