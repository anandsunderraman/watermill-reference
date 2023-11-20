package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"your-go-module-name/asyncapi"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/pkg/errors"
)

func main() {

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	err := startKafkaPublishers(ctx)
  if err != nil {
    log.Fatalf("error starting kafka publishers: %s", err)
	}

	router, err := asyncapi.GetRouter()
  if err != nil {
    log.Fatalf("error getting router: %s", err)
	}

	err = startKafkaSubscribers(ctx, router)
  if err != nil {
    log.Fatalf("error starting kafka subscribers: %s", err)
	}

	if err = router.Run(ctx); err != nil {
    log.Fatalf("error running watermill router: %s", err)
  }

}

func startKafkaPublishers(ctx context.Context) error {
	pub, err := asyncapi.GetKafkaPublisher(asyncapi.GetKafkaBroker())
	if err != nil {
		return errors.Wrap(err, "getting kafka publisher")
	}

  for i := 0; i < 5; i++ {
    l := asyncapi.LightMeasured {
      Id: 22,
      Lumens: 44,
      SentAt: time.Now().String(),
		}

		m, err := asyncapi.PayloadToMessage(l)
		if err != nil {
			return errors.Wrapf(err, "converting payload: %+v to message: %+v", l, err)
		}

		pub.Publish("light-measured", m)
    if err != nil {
      return errors.Wrapf(err, "publishing to topic: %s", "light-measured")
    }
    log.Printf("published message: %+v\n", l)
  }
  return nil
}

func startKafkaSubscribers(ctx context.Context, router *message.Router) error {
  kafkaSubscriber, err := asyncapi.GetKafkaSubscriber(asyncapi.GetKafkaBroker())
  if err != nil {
    return err
  }

  asyncapi.ConfigureKafkaSubscriptionHandlers(router, kafkaSubscriber)
  return nil
}
