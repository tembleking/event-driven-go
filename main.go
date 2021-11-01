package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/tembleking/event-driven-go/pkg/domain"
	"github.com/tembleking/event-driven-go/pkg/domain/consumer"
	"github.com/tembleking/event-driven-go/pkg/domain/producer"
)

func main() {
	timeoutCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	ctx := signalInterruptor(timeoutCtx)

	messageBroker := domain.NewBroker()
	go messageBroker.Start(ctx)

	aProducer := producer.NewProducer(messageBroker)
	aConsumer := consumer.NewConsumer(messageBroker)

	go aConsumer.Work(ctx)
	go startProducing(aProducer)

	<-ctx.Done()
	log.Println("quitting")
}

func startProducing(aProducer *producer.Producer) {
	for {
		time.Sleep(time.Second)
		aProducer.Produce()
	}
}

func signalInterruptor(parent context.Context) context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancelFunc := context.WithCancel(parent)
	go func() {
		aSignal := <-c
		log.Printf("cancelling from signal [signal=%s]", aSignal.String())
		cancelFunc()
	}()

	return ctx
}
