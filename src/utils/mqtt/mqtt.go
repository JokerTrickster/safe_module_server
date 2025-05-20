package mqtt

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/eclipse/paho.golang/paho"
)

var (
	client      *paho.Client
	router      *paho.StandardRouter
	respWaiters sync.Map // map[string]chan *paho.Publish
)

// MQTTInit initializes MQTT client
func MQTTInit() error {
	// MQTT 브로커에 연결
	conn, err := net.Dial("tcp", "192.168.0.6:1883")
	if err != nil {
		return fmt.Errorf("failed to connect to broker: %v", err)
	}

	router = paho.NewStandardRouter()
	client = paho.NewClient(paho.ClientConfig{
		Conn:   conn,
		Router: router,
	})
	clientID := fmt.Sprintf("go-mqtt-logan2-%d", time.Now().UnixNano())
	// CONNECT 패킷 전송
	_, err = client.Connect(context.Background(), &paho.Connect{
		ClientID: clientID,
	})
	if err != nil {
		return fmt.Errorf("failed to connect: %v", err)
	}

	Subscribe("/sensor/datas", 2, SensorDataHandler)

	return nil
}

// Subscribe subscribes to a topic
func Subscribe(topic string, qos byte, handler func(*paho.Publish)) error {
	if router == nil {
		return fmt.Errorf("router not initialized")
	}
	router.RegisterHandler(topic, handler)

	_, err := client.Subscribe(context.Background(), &paho.Subscribe{
		Subscriptions: []paho.SubscribeOptions{
			{
				Topic: topic,
				QoS:   qos,
			},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to subscribe: %v", err)
	}

	fmt.Printf("구독 성공 토픽 : %s\n", topic)
	return nil
}

// Publish publishes a message to a topic
func Publish(topic string, qos byte, payload interface{}, correlationData string, responseTopic string) error {
	var message []byte
	switch v := payload.(type) {
	case string:
		message = []byte(v)
	case []byte:
		message = v
	default:
		return fmt.Errorf("unsupported payload type")
	}

	_, err := client.Publish(context.Background(), &paho.Publish{
		Topic:   topic,
		QoS:     qos,
		Payload: message,
		Properties: &paho.PublishProperties{
			CorrelationData: []byte(correlationData),
			ResponseTopic:   responseTopic,
			ContentType:     "application/json",
		},
	})
	if err != nil {
		return fmt.Errorf("failed to publish: %v", err)
	}

	fmt.Printf("Published to topic: %s\n", topic)
	return nil
}

// Close closes the MQTT connection
func Close() {
	if client != nil {
		client.Disconnect(&paho.Disconnect{})
	}
}

// Publish & wait for response (timeout 지원)
func PublishAndWaitForResponse(topic string, qos byte, payload interface{}, correlationID string, responseTopic string, timeout time.Duration) (*paho.Publish, error) {
	// 2. 응답 채널 준비 & 등록
	respCh := make(chan *paho.Publish, 1)
	respWaiters.Store(correlationID, respCh)
	defer respWaiters.Delete(correlationID)

	// 3. Publish 요청 전송
	err := Publish(topic, qos, payload, correlationID, responseTopic)
	if err != nil {
		return nil, err
	}

	// 4. 응답 대기 (timeout)
	select {
	case resp := <-respCh:
		return resp, nil
	case <-time.After(timeout):
		return nil, fmt.Errorf("timeout waiting for response")
	}
}
