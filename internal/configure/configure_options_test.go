package configure

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestSetup_OptionsFlagEnvPrecedence(t *testing.T) {
	sb := NewSetupBlockletImpl(nil)
	cmd := cobra.Command{Use: "test"}

	// set env vars
	t.Setenv("label", "envlbl")
	t.Setenv("format", "envfmt")
	t.Setenv("color", "#abc")
	t.Setenv("bgcolor", "#def")

	sb.Setup(&cmd)
	// flags should override env when set
	cmd.PersistentFlags().Set("label", "flaglbl")
	cmd.PersistentFlags().Set("format", "flagfmt")
	cmd.PersistentFlags().Set("color", "#111")
	cmd.PersistentFlags().Set("bgcolor", "#222")

	opts := sb.Options()
	if opts.Label() != "flaglbl" {
		t.Fatalf("label mismatch: %q", opts.Label())
	}
	if opts.Format() != "flagfmt" {
		t.Fatalf("format mismatch: %q", opts.Format())
	}
	if opts.ForegroundColor() != "#111" {
		t.Fatalf("fg mismatch: %q", opts.ForegroundColor())
	}
	if opts.BackgroundColor() != "#222" {
		t.Fatalf("bg mismatch: %q", opts.BackgroundColor())
	}
}
