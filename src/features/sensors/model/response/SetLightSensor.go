package response

type ResSetLightSensor struct {
	SensorID    string `json:"sensorID"`
	LightStatus string `json:"lightStatus"`
}
