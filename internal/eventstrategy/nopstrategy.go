package eventstrategy

import (
	"github.com/spetix/bar-out-adapters/internal/events"
	"github.com/spetix/bar-out-adapters/pkg/barout/models"
)

type NopStrategy struct{}

func NewNopStrategy() *events.EventManagerImpl {
	em := events.NewEventManagerImpl()
	em.SetStrategy(&NopStrategy{})
	return em
}

func (e *NopStrategy) Register(k models.Event, h models.EventHdlr) {
	// Do nothing
}

func (e *NopStrategy) Execute() error {
	// Do nothing
	return nil
}
