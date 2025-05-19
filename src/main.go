package main

import (
	"main/features/sensors/handler"
	"main/utils/db"
	_log "main/utils/log"
	"main/utils/mqtt"

	swaggerDocs "main/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

//export PATH=$PATH:~/go/bin
func main() {
	// Echo 인스턴스 생성
	e := echo.New()

	// 로그 미들웨어 설정
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	if err := db.InitMongoDB(); err != nil {
		return
	}
	if err := _log.InitLogger(); err != nil {
		return
	}

	if err := mqtt.MQTTInit(); err != nil {
		return
	}

	if err := handler.NewSensorHandler(e); err != nil {
		return
	}

	// Swagger UI
	swaggerDocs.SwaggerInfo.Host = "localhost:8080"
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// 서버 시작
	e.Logger.Info("Server is starting on port 8080...")
	e.Logger.Fatal(e.Start(":8080"))
	return
}
