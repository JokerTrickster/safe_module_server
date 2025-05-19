package mqtt

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTTClient struct {
	client mqtt.Client
}

var MqttClient *MQTTClient

func MQTTInit() error {
	var err error
	MqttClient, err = NewMQTTClient("tcp://192.168.0.6:1883", "go-mqtt-client")
	if err != nil {
		fmt.Println("MQTT 클라이언트 생성 실패")
		defer MqttClient.Disconnect()
		return err
	}

	if err := MQTTTopicnIit(); err != nil {
		return err
	}

	return nil
}

// NewMQTTClient creates a new MQTT client instance
func NewMQTTClient(broker string, clientID string) (*MQTTClient, error) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientID)
	opts.SetProtocolVersion(5)
	opts.SetKeepAlive(60 * time.Second)
	opts.SetPingTimeout(1 * time.Second)
	opts.SetCleanSession(true)
	opts.SetAutoReconnect(true)
	opts.SetConnectionLostHandler(func(client mqtt.Client, err error) {
		log.Printf("Connection lost: %v", err)
	})
	opts.SetOnConnectHandler(func(client mqtt.Client) {
		log.Println("Connected to MQTT broker")
	})

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, fmt.Errorf("failed to connect to MQTT broker: %v", token.Error())
	}

	return &MQTTClient{client: client}, nil
}

// 토픽 구독 등록
func MQTTTopicnIit() error {
	err := MQTTSubscribeRegister("/sensor/datas", 2, SensorDataHandler)
	if err != nil {
		return err
	}

	err = MQTTSubscribeRegister("/sensor/light", 2, LightDataHandler)
	if err != nil {
		return err
	}
	return nil
}

// 토픽 구독 등록
func MQTTSubscribeRegister(topic string, qos byte, callback mqtt.MessageHandler) error {
	err := MqttClient.Subscribe(topic, qos, callback)
	if err != nil {
		return err
	}
	fmt.Println("MQTT 토픽 구독 성공 ", topic, qos)
	return nil
}

// 토픽 메시지 발행
func MQTTMessagePublish(topic string, qos byte, retained bool, payload interface{}) error {
	err := MqttClient.Publish(topic, qos, retained, payload)
	if err != nil {
		return err
	}
	return nil
}

// Publish sends a message to a topic
func (m *MQTTClient) Publish(topic string, qos byte, retained bool, payload interface{}) error {
	token := m.client.Publish(topic, qos, retained, payload)
	if token.Wait() && token.Error() != nil {
		return fmt.Errorf("failed to publish message: %v", token.Error())
	}
	return nil
}

// Subscribe subscribes to a topic
func (m *MQTTClient) Subscribe(topic string, qos byte, callback mqtt.MessageHandler) error {
	token := m.client.Subscribe(topic, qos, callback)
	if token.Wait() && token.Error() != nil {
		return fmt.Errorf("failed to subscribe to topic: %v", token.Error())
	}
	return nil
}

// Unsubscribe unsubscribes from a topic
func (m *MQTTClient) Unsubscribe(topic string) error {
	token := m.client.Unsubscribe(topic)
	if token.Wait() && token.Error() != nil {
		return fmt.Errorf("failed to unsubscribe from topic: %v", token.Error())
	}
	return nil
}

// Disconnect disconnects from the MQTT broker
func (m *MQTTClient) Disconnect() {
	m.client.Disconnect(250)
}
