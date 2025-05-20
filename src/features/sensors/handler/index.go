package handler

import (
	"main/features/sensors/repository"
	"main/features/sensors/usecase"
	"time"

	"main/utils/db"

	"github.com/labstack/echo/v4"
)

func NewSensorHandler(c *echo.Echo) error {
	NewGetSensorHandler(c, usecase.NewGetSensorUseCase(repository.NewGetSensorRepository(db.Client), 10*time.Second))
	NewSetLightSensorHandler(c, usecase.NewSetLightSensorUseCase(repository.NewSetLightSensorRepository(db.Client), 10*time.Second))
	NewTopicRegisterSensorHandler(c, usecase.NewTopicRegisterSensorUseCase(repository.NewTopicRegisterSensorRepository(db.Client), 10*time.Second))

	return nil
}
