package request

type ReqConfirmEventSensor struct {
	SensorID string `json:"sensor_id" validate:"required"`
	Type     string `json:"type" validate:"required"`
	Status   string `json:"status" validate:"required"`
}
