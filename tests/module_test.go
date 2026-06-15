package tests

import (
	"github.com/spetix/bar-out-adapters/generated/mocks"
	"github.com/spetix/bar-out-adapters/pkg/barout"
	"github.com/spf13/cobra"
)

func ExampleNewSetupBlocklet_ok() {
	m := new(mocks.MockData)
	m.On("Short").Return("name")
	m.On("Long").Return("details")
	m.On("BackgroundColor").Return("background")
	m.On("ForegroundColor").Return("foreground")
	m.On("Label").Return("name")

	sb := barout.NewSetupBlocklet(nil)
	cmd := cobra.Command{
		Use: "test",
	}
	// Force protocol to json via flag so Setup picks it up
	sb.Setup(&cmd)
	cmd.PersistentFlags().Set("protocol", "json")
	out := sb.GetOutput()
	out.Print(m)
	// Output:
	// {"text":"name","tooltip":"details","alt":"name","background-color":"background","foreground-color":"foreground"}
}
