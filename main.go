package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"your-go-module-name/asyncapi"

	"github.com/ThreeDotsLabs/watermill/message"
)

func main() {

  ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
  defer stop()

  err := startAMQPPublishers(ctx)
  if err != nil {
    log.Fatalf("error starting amqp publishers: %s", err)
  }

  router, err := asyncapi.GetRouter()
  if err != nil {
    log.Fatalf("error getting router: %s", err)
  }

  err = startAMQPSubscribers(ctx, router)
  if err != nil {
    log.Fatalf("error starting amqp subscribers: %s", err)
  }

  if err = router.Run(ctx); err != nil {
    log.Fatalf("error running watermill router: %s", err)
  }
}

func startAMQPSubscribers(ctx context.Context, router *message.Router) error {
  amqpSubscriber, err := asyncapi.GetAMQPSubscriber(asyncapi.GetAMQPURI())
  if err != nil {
    return err
  }

  asyncapi.ConfigureAMQPSubscriptionHandlers(router, amqpSubscriber)
  return nil
}

func startAMQPPublishers(ctx context.Context) error {
  pub, err := asyncapi.GetAMQPPublisher(asyncapi.GetAMQPURI())
  for i := 0; i < 5; i++ {
    l := asyncapi.LightMeasured {
      Id: 22,
      Lumens: 44,
      SentAt: "2022-02-01",
    }


    err = asyncapi.PublishToAMQP(ctx, pub, "light/measured", l)
    if err != nil {
      return err
    }
    log.Println("published message")
  }
  return nil
}

