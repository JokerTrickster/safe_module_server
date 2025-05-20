package mqtt

import (
	"fmt"

	"github.com/eclipse/paho.golang/paho"
)

// 응답 핸들러 등록
func SensorLightResponseHandler(p *paho.Publish) {
	fmt.Println("SetLightResponseHandler 호출", string(p.Payload))
	if p.Properties == nil || p.Properties.CorrelationData == nil {
		return
	}
	corrID := string(p.Properties.CorrelationData)
	if chVal, ok := respWaiters.Load(corrID); ok {
		ch := chVal.(chan *paho.Publish)
		ch <- p // 응답 전송
	}
}

// 조명 정보 가져오는 핸들러
func GetLightResponseHandler(p *paho.Publish) {
	fmt.Println("GetLightResponseHandler 호출", string(p.Payload))
	if p.Properties == nil || p.Properties.CorrelationData == nil {
		return
	}
	corrID := string(p.Properties.CorrelationData)
	if chVal, ok := respWaiters.Load(corrID); ok {
		ch := chVal.(chan *paho.Publish)
		ch <- p // 응답 전송
	}
}
