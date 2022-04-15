package asyncapi

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
)

// GetKafkaPublisher returns an amqp publisher based on the URI
func GetKafkaPublisher(kafkaBrokers string) (*kafka.Publisher, error) {
	kafkaConfig := kafka.PublisherConfig{
		Brokers:   []string{kafkaBrokers},
		Marshaler: kafka.DefaultMarshaler{},
	}

	return kafka.NewPublisher(
		kafkaConfig,
		watermill.NewStdLogger(false, false),
	)
}
