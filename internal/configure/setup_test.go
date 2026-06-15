package configure

import (
	"testing"

	"github.com/spf13/cobra"

	"github.com/spetix/bar-out-adapters/internal/protocols"
)

func TestSetup_GetOutputSelection(t *testing.T) {
	sb := NewSetupBlockletImpl(nil)
	cmd := cobra.Command{Use: "test"}
	sb.Setup(&cmd)

	// json
	cmd.PersistentFlags().Set("protocol", "json")
	out := sb.GetOutput()
	if _, ok := out.(*protocols.JsonOut); !ok {
		t.Fatalf("expected JsonOut for json protocol, got %T", out)
	}

	// raw
	cmd.PersistentFlags().Set("protocol", "raw")
	out = sb.GetOutput()
	if _, ok := out.(*protocols.RawOut); !ok {
		t.Fatalf("expected RawOut for raw protocol, got %T", out)
	}

	// waybar
	cmd.PersistentFlags().Set("protocol", "waybar")
	out = sb.GetOutput()
	if _, ok := out.(*protocols.WaybarOut); !ok {
		t.Fatalf("expected WaybarOut for waybar protocol, got %T", out)
	}
}
func TestSetup_GetOutputDefaultBranch(t *testing.T) {
	sb := NewSetupBlockletImpl(nil)
	sb.protocol = protocols.BlockletProtocol("invalid")
	out := sb.GetOutput()
	if _, ok := out.(*protocols.RawOut); !ok {
		t.Fatalf("expected RawOut for unknown protocol default, got %T", out)
	}
}
