package eventstrategy

import (
	"github.com/spetix/bar-out-adapters/internal/events"
	"github.com/spetix/bar-out-adapters/pkg/barout/models"
)

type commonStrategy struct {
	eventManager *events.EventManagerImpl
}

type strategyBuilder func(c *commonStrategy) models.EventStrategy

func NewStrategy(sf strategyBuilder) *events.EventManagerImpl {
	em := events.NewEventManagerImpl()
	st := sf(&commonStrategy{
		eventManager: em,
	})
	em.SetStrategy(st)
	return em
}
