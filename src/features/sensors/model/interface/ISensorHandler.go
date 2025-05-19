package _interface

import (
	"github.com/labstack/echo/v4"
)

type IGetSensorHandler interface {
	GetSensor(c echo.Context) error
}
