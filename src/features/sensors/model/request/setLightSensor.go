package request

type ReqSetLightSensor struct {
	SensorID string `json:"sensorID" example:"30:ED:A0:BA:13:20"`
	Status   string `json:"status" example:"on / off"` // true 가 On false Off
}
