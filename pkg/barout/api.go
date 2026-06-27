package barout

import (
	"os"

	"github.com/spetix/bar-out-adapters/internal/configure"
	"github.com/spetix/bar-out-adapters/internal/events"
	"github.com/spetix/bar-out-adapters/pkg/barout/models"
)

// NewSetupBlocket selects the right driver for the output
func NewSetupBlocklet(f models.Formatter) models.SetupBlocklet {
	return configure.NewSetupBlockletImpl(f)
}

// NewEvent creates a new object of type Event
func NewEvent(et models.EventType) models.Event {
	return events.NewEventImpl(et)
}

// NewSignal creates a new Event object for signals
func NewSignal(s os.Signal) models.Event {
	return events.NewEventImplSignal(s)
}
