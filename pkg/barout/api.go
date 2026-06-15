package barout

import (
	"github.com/spetix/bar-out-adapters/internal/configure"
	"github.com/spetix/bar-out-adapters/pkg/barout/models"
)

// NewSetupBlocket selects the right driver for the output
func NewSetupBlocklet(f models.Formatter) models.SetupBlocklet {
	return configure.NewSetupBlockletImpl(f)
}
