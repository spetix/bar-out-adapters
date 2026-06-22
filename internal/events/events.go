package events

import (
	"github.com/spetix/bar-out-adapters/pkg/barout/models"
)

type EventImpl struct {
	name   models.EventType
	signal int
}

func (e *EventImpl) Name() models.EventType {
	return e.name
}

func (e *EventImpl) Signal() int {
	return e.signal
}

type EventManagerImpl struct {
	handlers models.EventHandlers
	strategy models.EventStrategy
}

func NewEventManagerImpl() *EventManagerImpl {
	return &EventManagerImpl{
		handlers: make(models.EventHandlers, 0),
	}
}

func (e *EventManagerImpl) Register(ev models.EventType, hdlr models.EventHdlr) {
	if _, exists := e.handlers[ev]; exists {
		return
	}
	e.handlers[ev] = hdlr
}

func (e *EventManagerImpl) Handlers() models.EventHandlers {
	return e.handlers
}

func (e *EventManagerImpl) SetStrategy(strategy models.EventStrategy) {
	e.strategy = strategy
}

func (e *EventManagerImpl) Run() {
	if e.strategy != nil {
		e.strategy.Execute()
	}
}
