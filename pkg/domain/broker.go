package domain

import (
	"context"
	"log"
)

type Broker struct {
	bufferSizeBeforeEventsAreDropped int
	publishCh                        chan Event
	subscribeCh                      chan chan Event
	unsubscribeCh                    chan chan Event
}

func (b *Broker) Start(ctx context.Context) {
	subs := map[chan Event]struct{}{}
	for {
		select {
		case msgCh := <-b.subscribeCh:
			subs[msgCh] = struct{}{}
		case msgCh := <-b.unsubscribeCh:
			delete(subs, msgCh)
		case event := <-b.publishCh:
			for subCh := range subs {
				select {
				case subCh <- event:
				default: // protects the broker by dropping events
					log.Printf("broker: warning, event dropped [type=%T, ID=%s, happenedOn=%s]", event, event.ID(), event.HappenedOn())
				}
			}
		case <-ctx.Done():
			return
		}
	}
}

func (b *Broker) Subscribe() chan Event {
	msgCh := make(chan Event, b.bufferSizeBeforeEventsAreDropped)
	b.subscribeCh <- msgCh
	return msgCh
}

func (b *Broker) Unsubscribe(msgCh chan Event) {
	b.unsubscribeCh <- msgCh
}

func (b *Broker) Publish(msg Event) {
	b.publishCh <- msg
}

func NewBroker() *Broker {
	bufferSizeBeforeEventsAreDropped := 3
	return NewBrokerWithBufferSize(bufferSizeBeforeEventsAreDropped)
}

func NewBrokerWithBufferSize(bufferSizeBeforeEventsAreDropped int) *Broker {
	return &Broker{
		bufferSizeBeforeEventsAreDropped: bufferSizeBeforeEventsAreDropped,
		publishCh:                        make(chan Event),
		subscribeCh:                      make(chan chan Event),
		unsubscribeCh:                    make(chan chan Event),
	}
}
