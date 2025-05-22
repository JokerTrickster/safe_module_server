package request

type ReqSetPositionSensor struct {
	SensorID string   `json:"sensorID" validate:"required"`
	Position Position `json:"position" validate:"required"`
}

type Position struct {
	X float64 `json:"x" validate:"required"`
	Y float64 `json:"y" validate:"required"`
}
