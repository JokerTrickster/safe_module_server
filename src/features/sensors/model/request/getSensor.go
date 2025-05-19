package request

type ReqGetSensor struct {
	SensorID string `query:"sensorID" validate:"required"`
}
