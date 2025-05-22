package usecase

import (
	"main/features/sensors/model/request"
	"main/utils/db"
	"time"
)

func CreateThresholdDTO(req *request.ReqSetThresholdSensor) db.SensorThresholdDTO {
	now := time.Now()
	return db.SensorThresholdDTO{
		Name:      req.Name,
		Unit:      req.Unit,
		Threshold: req.Threshold,
		CreatedAt: &now,
		DeletedAt: nil,
		UpdatedAt: &now,
	}
}
