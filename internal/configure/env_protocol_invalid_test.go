package configure

import (
	"os"
	"testing"

	"github.com/spetix/bar-out-adapters/internal/protocols"
	"github.com/spf13/cobra"
)

func TestSetup_InvalidEnvProtocolDefaultsToRaw(t *testing.T) {
	sb := NewSetupBlockletImpl(nil)
	cmd := cobra.Command{Use: "test"}

	os.Setenv("protocol", "invalid")
	defer os.Unsetenv("protocol")

	sb.Setup(&cmd)
	out := sb.GetOutput()
	if _, ok := out.(*protocols.RawOut); !ok {
		t.Fatalf("expected RawOut when env protocol invalid, got %T", out)
	}
}
