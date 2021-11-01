package producer

import (
	"math/rand"
	"strconv"
	"time"
)

type ProducedSomethingEvent struct {
	id         string
	happenedOn time.Time
	produced   int
}

func (e *ProducedSomethingEvent) ID() string {
	return e.id
}

func (e *ProducedSomethingEvent) HappenedOn() time.Time {
	return e.happenedOn
}

func (e *ProducedSomethingEvent) Produced() int {
	// using getter method to avoid event being modified
	return e.produced
}

func NewProducedSomethingEvent(produced int) *ProducedSomethingEvent {
	return &ProducedSomethingEvent{
		id:         strconv.Itoa(rand.Int()),
		happenedOn: time.Now(),
		produced:   produced,
	}
}
