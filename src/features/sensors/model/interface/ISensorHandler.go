package _interface

import (
	"github.com/labstack/echo/v4"
)

type IGetSensorHandler interface {
	GetSensor(c echo.Context) error
}

type ISetLightSensorHandler interface {
	SetLightSensor(c echo.Context) error
}

type ITopicRegisterSensorHandler interface {
	TopicRegisterSensor(c echo.Context) error
}

type IGetLightSensorHandler interface {
	GetLightSensor(c echo.Context) error
}
