package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"your-go-module-name/asyncapi"
)

func main() {

  for i := 0; i < 5; i++ {
    l := asyncapi.LightMeasured {
      Id: 22,
      Lumens: 44,
      SentAt: "2022-02-01",
    }
    pub, err := asyncapi.GetAMQPPublisher(asyncapi.GetAMQPURI())

    err = asyncapi.PublishToAMQP(pub, "light/measured", l)
    if err != nil {
      log.Fatalf("error publishing payload: %+v error: %s", l, err)
    }
  }

  router, err := asyncapi.GetRouter()
  if err != nil {
    log.Fatalf("error creating watermill router: %s", err)
  }


  amqpSubscriber, err := asyncapi.GetAMQPSubscriber(asyncapi.GetAMQPURI())
  if err != nil {
    log.Fatalf("error creating amqpSubscriber: %s", err)
  }

  asyncapi.ConfigureAMQPSubscriptionHandlers(router, amqpSubscriber)


  ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
  defer stop()
  if err = router.Run(ctx); err != nil {
    log.Fatalf("error running watermill router: %s", err)
  }


}

