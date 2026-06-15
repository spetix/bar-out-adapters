package barout_test

import (
	"testing"

	"github.com/spf13/cobra"

	"github.com/spetix/bar-out-adapters/internal/protocols"
	"github.com/spetix/bar-out-adapters/pkg/barout"
)

func TestNewSetupBlocklet_DefaultOptions(t *testing.T) {
	sb := barout.NewSetupBlocklet(nil)
	cmd := cobra.Command{Use: "test"}
	sb.Setup(&cmd)
	opts := sb.Options()
	if opts.Label() != "🎄" {
		t.Fatalf("unexpected default label: %q", opts.Label())
	}
	if opts.ForegroundColor() != "#ffffff" {
		t.Fatalf("unexpected default fg: %q", opts.ForegroundColor())
	}
	if opts.BackgroundColor() != "#000000" {
		t.Fatalf("unexpected default bg: %q", opts.BackgroundColor())
	}
}

func TestNewSetupBlocklet_GetOutputSelection(t *testing.T) {
	sb := barout.NewSetupBlocklet(nil)
	cmd := cobra.Command{Use: "test"}
	sb.Setup(&cmd)

	cmd.PersistentFlags().Set("protocol", "json")
	out := sb.GetOutput()
	if _, ok := out.(*protocols.JsonOut); !ok {
		t.Fatalf("expected JsonOut for json protocol, got %T", out)
	}

	cmd.PersistentFlags().Set("protocol", "raw")
	out = sb.GetOutput()
	if _, ok := out.(*protocols.RawOut); !ok {
		t.Fatalf("expected RawOut for raw protocol, got %T", out)
	}

	cmd.PersistentFlags().Set("protocol", "i3blocks")
	out = sb.GetOutput()
	if _, ok := out.(*protocols.RawOut); !ok {
		t.Fatalf("expected RawOut for i3blocks protocol, got %T", out)
	}

	cmd.PersistentFlags().Set("protocol", "waybar")
	out = sb.GetOutput()
	if _, ok := out.(*protocols.WaybarOut); !ok {
		t.Fatalf("expected WaybarOut for waybar protocol, got %T", out)
	}
}
