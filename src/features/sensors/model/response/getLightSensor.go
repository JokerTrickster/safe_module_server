package response

type ResGetLightSensor struct {
	Status string `json:"status" example:"on / off"`
}

type LightResponse struct {
	Status string `json:"status"`
}
