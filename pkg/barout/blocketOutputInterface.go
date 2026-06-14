package barout

import (
	"fmt"

	"github.com/spetix/bar-out-adapters/pkg/barout/data"
)

// BlockletOutput is the interface that wraps the Print method. It is used to output data in different formats.
// An implementation of this interface is provided for each supported protocol.

type BlockletOutput interface {
	Print(d data.Data)
}

type BlockletProtocol string

const (
	ProtocolRaw      BlockletProtocol = "raw"
	ProtocolI3Blocks BlockletProtocol = "i3blocks"
	ProtocolWaybar   BlockletProtocol = "waybar"
	ProtocolJson     BlockletProtocol = "json"
)

func (p BlockletProtocol) String() string {
	return string(p)
}

func (p BlockletProtocol) Type() string {
	return "BlockletProtocol"
}

func (p *BlockletProtocol) Set(protoStr string) error {
	switch protoStr {
	case "raw":
		*p = ProtocolRaw
	case "i3blocks":
		*p = ProtocolI3Blocks
	case "waybar":
		*p = ProtocolWaybar
	case "json":
		*p = ProtocolJson
	default:
		*p = ProtocolRaw
		return fmt.Errorf("Invalid protocol %s", protoStr)
	}
	return nil
}
