package response

type ResListSensor struct {
	SensorList []ListSensor `json:"sensorList"`
}

type ListSensor struct {
	SensorID    string   `json:"sensorID"`
	LightStatus string   `json:"lightStatus"`
	Sensors     []Sensor `json:"sensors"`
}
