package mqtt

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func LightDataHandler(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message on topic!! %s: %s\n", msg.Topic(), string(msg.Payload()))
}
