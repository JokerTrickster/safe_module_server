package request

type ReqSetThresholdSensor struct {
	Name      string  `json:"name" example:"co2"`
	Threshold float64 `json:"threshold" example:"3000"`
	Unit      string  `json:"unit" example:"ppm"`
}
