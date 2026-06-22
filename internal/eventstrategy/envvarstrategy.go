package eventstrategy

import (
	"os"

	"github.com/spetix/bar-out-adapters/internal/events"
	"github.com/spetix/bar-out-adapters/pkg/barout/models"
)

type EnvVarStrategy struct {
	hdlrs models.EventHandlers
}

func NewEnvVarStrategy(e *events.EventManagerImpl) {
	e.SetStrategy(&EnvVarStrategy{
		hdlrs: e.Handlers(),
	})
}

func (e *EnvVarStrategy) Execute() {
	if ev := os.Getenv("BUTTON_BLOCK"); ev != "" {
		switch ev {
		case models.LeftButton.String():
			if hdlr, exists := e.hdlrs[models.LeftButton]; exists {
				hdlr()
			}
		case models.MiddleButton.String():
			if hdlr, exists := e.hdlrs[models.MiddleButton]; exists {
				hdlr()
			}
		case models.RightButton.String():
			if hdlr, exists := e.hdlrs[models.RightButton]; exists {
				hdlr()
			}
		case models.ScrollUp.String():
			if hdlr, exists := e.hdlrs[models.ScrollUp]; exists {
				hdlr()
			}
		case models.ScrollDown.String():
			if hdlr, exists := e.hdlrs[models.ScrollDown]; exists {
				hdlr()
			}
		default:
			return
		}
	}
}
