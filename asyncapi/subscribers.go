package asyncapi

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
)

// GetKafkaSubscriber returns an amqp subscriber based on the URI
func GetKafkaSubscriber(kafkaBrokers string) (*kafka.Subscriber, error) {

	cfg := kafka.DefaultSaramaSubscriberConfig()
	// cfg.Consumer.Offsets.Initial = sarama.OffsetOldest
	kafkaConfig := kafka.SubscriberConfig{
		Brokers:               []string{kafkaBrokers},
		Unmarshaler:           kafka.DefaultMarshaler{},
		OverwriteSaramaConfig: cfg,
		ConsumerGroup:         "test_consumer_group",
	}

	return kafka.NewSubscriber(
		kafkaConfig,
		watermill.NewStdLogger(false, false),
	)
}
