package barout_test

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/spf13/cobra"

	"github.com/spetix/bar-out-adapters/internal/protocols"
	"github.com/spetix/bar-out-adapters/pkg/barout"
	"github.com/spetix/bar-out-adapters/pkg/barout/models"
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

func TestNewSetupBlocklet_EventManager(t *testing.T) {
	sb := barout.NewSetupBlocklet(nil)
	cmd := cobra.Command{Use: "test"}
	cmd.PersistentFlags().Set("protocol", protocols.ProtocolI3Blocks.String())

	sb.Setup(&cmd)
	em := sb.EventManager()

	if em == nil {
		t.Fatal("expected non-nil EventManager")
	}
	em.Register(models.LeftButton, func() error {
		t.Log("Left Button Called")
		return nil
	})
	em.Register(models.RightButton, func() error {
		t.Fail()
		return fmt.Errorf("Right button not called")
	})
	oldVal := os.Getenv("BUTTON_BLOCK")
	if oldVal != "" {
		defer os.Setenv("BUTTON_BLOCK", oldVal)
	} else {
		defer os.Unsetenv("BUTTON_BLOCK")
	}
	os.Setenv("BUTTON_BLOCK", models.LeftButton.String())
	em.Run()
}

func TestNewSetupBlocklet_GetOutputSelection(t *testing.T) {
	sb := barout.NewSetupBlocklet(nil)
	cmd := cobra.Command{Use: "test"}
	sb.Setup(&cmd)

	tests := []struct {
		name     string
		protocol string
		wantType any
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
			em := sb.EventManager()
			if em == nil {
				t.Fatalf("expected non-nil EventManager for %s protocol", tt.protocol)
			}
		})
	}
}
