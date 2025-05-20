package request

type ReqGetLightSensor struct {
	SensorID string `query:"sensorID" validate:"required" example:"30:ED:A0:BA:13:20"`
}
