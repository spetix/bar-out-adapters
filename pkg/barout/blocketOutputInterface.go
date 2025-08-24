package barout

import (
	"os"

	"github.com/spetix/bar-out-adapters/internal/protocols"
)

type baseOutput struct {
	Device *os.File
}

type BlockletOutput interface {
	Print(d Data)
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
