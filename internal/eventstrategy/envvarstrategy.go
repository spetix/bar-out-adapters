package eventstrategy

import (
	"fmt"
	"os"

	"github.com/spetix/bar-out-adapters/internal/events"
	"github.com/spetix/bar-out-adapters/pkg/barout/models"
)

type EnvVarStrategy struct {
	*commonStrategy
}

func NewEnvVarStrategy(c *commonStrategy) models.EventStrategy {
	s := &EnvVarStrategy{}
	s.commonStrategy = c
	return s
}

func (e *EnvVarStrategy) Register(k models.Event, h models.EventHdlr) {
	registered := e.eventManager.Events()
	if _, exists := registered[k]; !exists {
		registered[k] = h
	}
}

func (e *EnvVarStrategy) Execute() error {
	if ev := os.Getenv("BUTTON_BLOCK"); ev != "" {
		registry := e.eventManager.Events()
		switch ev {
		case models.LeftButton.String():

			if hdlr, exists := registry[events.NewEventImpl(models.LeftButton)]; exists {
				return hdlr()
			}
		case models.MiddleButton.String():
			if hdlr, exists := registry[events.NewEventImpl(models.MiddleButton)]; exists {
				return hdlr()
			}
		case models.RightButton.String():
			if hdlr, exists := registry[events.NewEventImpl(models.RightButton)]; exists {
				return hdlr()
			}
		case models.ScrollUp.String():
			if hdlr, exists := registry[events.NewEventImpl(models.ScrollUp)]; exists {
				return hdlr()
			}
		case models.ScrollDown.String():
			if hdlr, exists := registry[events.NewEventImpl(models.ScrollDown)]; exists {
				return hdlr()
			}
		case models.Signal.String():
			return fmt.Errorf("Signal not supposed to be received here")
		default:
			return fmt.Errorf("unknown event: %s", ev)
		}
	}
	return nil
}
