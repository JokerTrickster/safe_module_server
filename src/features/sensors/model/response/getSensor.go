package response

type ResGetSensor struct {
	SensorID    string   `json:"sensorID"`
	LightStatus bool     `json:"lightStatus"`
	Sensors     []Sensor `json:"sensors"`
}

type Sensor struct {
	Name   string  `json:"name"`
	Value  float64 `json:"value"`
	Status string  `json:"status"`
}
