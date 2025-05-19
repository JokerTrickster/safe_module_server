package response

type ResSetLightSensor struct {
	SensorID    string `json:"sensorID"`
	LightStatus bool   `json:"lightStatus"`
}
