package consumer

import (
	"context"
	"log"
	"time"

	"github.com/tembleking/event-driven-go/pkg/domain"
	"github.com/tembleking/event-driven-go/pkg/domain/producer"
)

type Consumer struct {
	events <-chan domain.Event
}

func (c *Consumer) Work(ctx context.Context) {
	for {
		select {
		case event := <-c.events:
			c.handleEvent(event)
		case <-ctx.Done():
			return
		}
	}
}

func (c *Consumer) handleEvent(domainEvent domain.Event) {
	time.Sleep(2 * time.Second)

	switch event := domainEvent.(type) {
	case *producer.ProducedSomethingEvent:
		log.Printf("consumer: the producer created something [produced=%d]", event.Produced())
	}
}

func NewConsumer(messageBroker *domain.Broker) *Consumer {
	return &Consumer{events: messageBroker.Subscribe()}
}
