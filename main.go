package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"your-go-module-name/asyncapi"
)

func main() {

	for i := 0; i < 5; i++ {
		l := asyncapi.LightMeasured{
			Id:     i,
			Lumens: i,
			SentAt: fmt.Sprintf("day %d", i),
		}
		pub, err := asyncapi.GetKafkaPublisher(asyncapi.GetKafkaBroker())

		err = asyncapi.PublishToKafka(pub, "light-measured", l)
		if err != nil {
			log.Fatalf("error publishing payload: %+v error: %s", l, err)
		}
		log.Printf("published message: %+v", l)
	}

	router, err := asyncapi.GetRouter()
	if err != nil {
		log.Fatalf("error creating watermill router: %s", err)
	}

	kafkaSubscriber, err := asyncapi.GetKafkaSubscriber(asyncapi.GetKafkaBroker())
	if err != nil {
		log.Fatalf("error creating kafkaSubscriber: %s", err)
	}
	asyncapi.ConfigureKafkaSubscriptionHandlers(router, kafkaSubscriber)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()
	if err = router.Run(ctx); err != nil {
		log.Fatalf("error running watermill router: %s", err)
	}

}
