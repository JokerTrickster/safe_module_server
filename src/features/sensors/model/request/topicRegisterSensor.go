package request

type ReqTopicRegisterSensor struct {
	Topic string `json:"topic"`
	Qos   int    `json:"qos"`
}
