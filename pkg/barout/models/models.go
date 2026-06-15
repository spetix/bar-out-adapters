package models

import (
	"github.com/spetix/bar-out-adapters/pkg/barout/data"
	"github.com/spf13/cobra"
)

type RenderOptions interface {
	Label() string
	Format() string
	ForegroundColor() string
	BackgroundColor() string
}

type Formatter interface {
	Render(s ...string) string
	Validate(f string) bool
}

type SetupBlocklet interface {
	Setup(*cobra.Command)
	GetOutput() BlockletOutput
	Options() RenderOptions
}

type BlockletOutput interface {
	Print(data data.Data)
}
