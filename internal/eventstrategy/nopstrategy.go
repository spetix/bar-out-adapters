package eventstrategy

import (
	"github.com/spetix/bar-out-adapters/internal/events"
)

type NopStrategy struct{}

func NewNopStrategy(e *events.EventManagerImpl) {
	e.SetStrategy(&NopStrategy{})
}

func (e *NopStrategy) Execute() {
	// Do nothing
}
