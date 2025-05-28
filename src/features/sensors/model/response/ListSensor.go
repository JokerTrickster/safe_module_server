package response

type ResListSensor struct {
	SensorList []ListSensor `json:"sensorList"`
}

type ListSensor struct {
	SensorID        string   `json:"sensorID"`
	LightStatus     string   `json:"lightStatus"`
	FireDetector    string   `json:"fireDetector"`
	MotionDetection string   `json:"motionDetection"`
	Position        Position `json:"position"`
	Sensors         []Sensor `json:"sensors"`
}
