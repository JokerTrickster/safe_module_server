package request

type ReqTopicRegisterSensor struct {
	Topic string `json:"topic" example:"/control/light/response/set/30:ED:A0:BA:13:20"`
	Qos   int    `json:"qos" example:"2"`
	Type  string `json:"type" example:"lightSet / lightGet"`
}
