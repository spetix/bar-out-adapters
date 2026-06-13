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

type BlockletProtocol string

const (
	Raw      BlockletProtocol = "raw"
	I3Blocks BlockletProtocol = "i3blocks"
	Waybar   BlockletProtocol = "waybar"
	Json     BlockletProtocol = "json"
)

func (p BlockletProtocol) String() string {
	return string(p)
}

func ParseProtocol(protoStr string) BlockletProtocol {
	switch protoStr {
	case "raw":
		return Raw
	case "i3blocks":
		return I3Blocks
	case "waybar":
		return Waybar
	case "json":
		return Json
	default:
		return Raw
	}
}

// New creates a new BlockletOutput based on the provided protocol string.
// It supports "json", "raw", and "waybar" protocols. If an unsupported protocol is provided, it defaults to "raw".
func New(protocol BlockletProtocol) BlockletOutput {

	switch protocol {
	case Raw, I3Blocks:
		return protocols.NewRawOut(os.Stdout)
	case Waybar:
		return protocols.NewWaybarOut(os.Stdout)
	case Json:
		return protocols.NewJsonOut(os.Stdout)
	default:
		return protocols.NewRawOut(os.Stdout)
	}
}
