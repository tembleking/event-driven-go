package producer

import (
	"log"

	"github.com/tembleking/event-driven-go/pkg/domain"
)

type Producer struct {
	messageBroker       *domain.Broker
	lastElementProduced int
}

func (p *Producer) Produce() {
	somethingProduced := p.lastElementProduced
	p.lastElementProduced++
	log.Printf("producer: I have produced something [produced=%d]", somethingProduced)

	p.messageBroker.Publish(NewProducedSomethingEvent(somethingProduced))
}

func NewProducer(messageBroker *domain.Broker) *Producer {
	return &Producer{messageBroker: messageBroker}
}
