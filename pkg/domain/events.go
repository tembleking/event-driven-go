package domain

import (
	"time"
)

type Event interface {
	ID() string
	HappenedOn() time.Time
}
