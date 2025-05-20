package request

type ReqGetSensor struct {
	SensorID string `query:"sensorID" validate:"required" example:"30:ED:A0:BA:13:20"`
}
