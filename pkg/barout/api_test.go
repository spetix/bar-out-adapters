package barout_test

import (
	"reflect"
	"testing"

	"github.com/spf13/cobra"

	"github.com/spetix/bar-out-adapters/internal/protocols"
	"github.com/spetix/bar-out-adapters/pkg/barout"
)

func TestBlockletProtocolSet(t *testing.T) {
	var proto protocols.BlockletProtocol
	if err := proto.Set("waybar"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if proto != protocols.ProtocolWaybar {
		t.Fatalf("expected protocol %q, got %q", protocols.ProtocolWaybar, proto)
	}
}

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

	tests := []struct {
		name     string
		protocol string
		wantType interface{}
	}{
		{name: "json", protocol: "json", wantType: &protocols.JsonOut{}},
		{name: "raw", protocol: "raw", wantType: &protocols.RawOut{}},
		{name: "i3blocks", protocol: "i3blocks", wantType: &protocols.RawOut{}},
		{name: "waybar", protocol: "waybar", wantType: &protocols.WaybarOut{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd.PersistentFlags().Set("protocol", tt.protocol)
			out := sb.GetOutput()
			if reflect.TypeOf(out) != reflect.TypeOf(tt.wantType) {
				t.Fatalf("expected %T for %s protocol, got %T", tt.wantType, tt.protocol, out)
			}
		})
	}
}
