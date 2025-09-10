package barout

import (
	"os"

	"github.com/spetix/bar-out-adapters/internal/protocols"
	"github.com/spetix/bar-out-adapters/pkg/barout/data"
)

type BlockletOutput interface {
	Print(d data.Data)
}

func New(protocol string) BlockletOutput {

	switch protocol {
	case "json":
		return protocols.NewJsonOut(os.Stdout)
	case "raw":
		return protocols.NewRawOut(os.Stdout)
	case "waybar":
		return protocols.NewWaybarOut(os.Stdout)
	default:
		return protocols.NewRawOut(os.Stdout)
	}
}
