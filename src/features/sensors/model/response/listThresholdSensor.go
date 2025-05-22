package response

type ResListThresholdSensor struct {
	ThresholdList []Threshold `json:"thresholdList"`
}

type Threshold struct {
	Name      string `json:"name"`
	Unit      string `json:"unit"`
	Threshold int    `json:"threshold"`
}
