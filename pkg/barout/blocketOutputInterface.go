package barout

import (
	"os"

	"github.com/spetix/bar-out-adapters/internal/protocols"
	"github.com/spetix/bar-out-adapters/pkg/barout/data"
)


// BlockletOutput is the interface that wraps the Print method. It is used to output data in different formats. 
// An implementation of this interface is provided for each supported protocol.

type BlockletOutput interface {
	Print(d data.Data)
}


// New creates a new BlockletOutput based on the provided protocol string.
// It supports "json", "raw", and "waybar" protocols. If an unsupported protocol is provided, it defaults to "raw".
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
