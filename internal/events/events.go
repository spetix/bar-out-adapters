package events

import (
	"os"

	"github.com/spetix/bar-out-adapters/pkg/barout/models"
)

type EventImpl struct {
	evType models.EventType
	signal os.Signal
}

func NewEventImpl(evType models.EventType) *EventImpl {
	return &EventImpl{
		evType: evType,
		signal: nil,
	}
}

func NewEventImplSignal(signal os.Signal) *EventImpl {
	return &EventImpl{
		evType: models.Signal,
		signal: signal,
	}
}

func (e *EventImpl) Type() models.EventType {
	return e.evType
}

func (e *EventImpl) Name() string {
	if e.IsSignal() {
		return e.signal.String()
	}
	return e.evType.String()
}

func (e *EventImpl) Signal() os.Signal {
	return e.signal
}

func (e *EventImpl) IsSignal() bool {
	return e.evType == models.Signal
}

type EventManagerImpl struct {
	strategy models.EventStrategy
	handlers models.EventHandlers
}

func NewEventManagerImpl() *EventManagerImpl {
	return &EventManagerImpl{}
}

func (e *EventManagerImpl) Register(k models.Event, hdlr models.EventHdlr) {
	e.strategy.Register(k, hdlr)
}

func (e *EventManagerImpl) Handlers() models.EventHandlers {
	return e.handlers
}

func (e *EventManagerImpl) Events() models.EventHandlers {
	m := make(models.EventHandlers, 0)
	for k, v := range e.handlers {
		if !k.IsSignal() {
			m[k] = v
		}
	}
	return m
}

func (e *EventManagerImpl) Signals() models.EventHandlers {
	m := make(models.EventHandlers, 0)
	for k, v := range e.handlers {
		if k.IsSignal() {
			m[k] = v
		}
	}
	return m
}

func (e *EventManagerImpl) SetStrategy(strategy models.EventStrategy) {
	e.strategy = strategy
}

func (e *EventManagerImpl) Run() error {
	if e.strategy != nil {
		return e.strategy.Execute()
	}
	return nil
}
