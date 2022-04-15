package asyncapi

import (
	"encoding/json"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"log"

	"github.com/ThreeDotsLabs/watermill/message"
)

// OnLightMeasured subscription handler for light/measured.
func OnLightMeasured(msg *message.Message) error {
	log.Printf("received message payload: %s", string(msg.Payload))

	var lm LightMeasured
	err := json.Unmarshal(msg.Payload, &lm)
	if err != nil {
		log.Printf("error unmarshalling message: %s, err is: %s", msg.Payload, err)
	}
	return nil
}

func PublishToKafka(p *kafka.Publisher, dest string, l LightMeasured) error {

	m, err := l.ToMessage()
	if err != nil {
		log.Fatalf("error converting payload: %+v to message error: %s", l, err)
	}

	return p.Publish(dest, &m)
}
