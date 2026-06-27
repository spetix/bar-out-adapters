package eventstrategy

import (
	"os/signal"

	"os"

	"github.com/spetix/bar-out-adapters/pkg/barout/models"
)

type SignalStrategy struct {
	*commonStrategy
	signalChannel chan os.Signal
}

func NewSignalStrategy(c *commonStrategy) models.EventStrategy {
	s := &SignalStrategy{
		signalChannel: make(chan os.Signal, 1),
	}
	s.commonStrategy = c
	return s
}

func (s *SignalStrategy) Register(k models.Event, h models.EventHdlr) {
	registry := s.commonStrategy.eventManager.Handlers()
	if _, exists := registry[k]; !exists {
		registry[k] = h
		signal.Notify(s.signalChannel, k.Signal())
	}
}

func (e *SignalStrategy) Execute() error {

	if len(e.eventManager.Signals()) > 0 {
		var sig os.Signal
		e.signalChannel <- sig
		for k, fHdlr := range e.eventManager.Signals() {
			if sig == k.Signal() {
				return fHdlr()
			}
		}
	}

	return nil
}
