package asyncapi

import (
	"context"
	"encoding/json"
	"log"

	"github.com/ThreeDotsLabs/watermill-amqp/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
)

// OnLightMeasured subscription handler for light/measured.
func OnLightMeasured(msg *message.Message) error {
    log.Printf("received message payload: %s", string(msg.Payload))

    var lm LightMeasured
    err := json.Unmarshal(msg.Payload, &lm)
    if err != nil {
        return err
    }
    return nil
}

func PublishToAMQP(ctx context.Context, p *amqp.Publisher, dest string, l LightMeasured) error {

    m, err := PayloadToMessage(l)
    if err != nil {
        return err
    }

    return p.Publish(dest, &m)
}

