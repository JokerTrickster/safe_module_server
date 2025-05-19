package request

type ReqSetLightSensor struct {
	SensorID string `json:"sensorID"`
	Status   bool   `json:"status"` // true 가 On false Off
}
