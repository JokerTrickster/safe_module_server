package response

type ResGetSensor struct {
	SensorID     string   `json:"sensorID"`
	LightStatus  string   `json:"lightStatus"`
	FireDetector string   `json:"fireDetector"`
	Position     Position `json:"position"`
	Sensors      []Sensor `json:"sensors"`
}

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Sensor struct {
	Name   string  `json:"name"`
	Value  float64 `json:"value"`
	Status string  `json:"status"`
	Unit   string  `json:"unit"`
}
