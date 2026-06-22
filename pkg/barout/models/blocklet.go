package models

import (
	"github.com/spf13/cobra"
)

type SetupBlocklet interface {
	Setup(*cobra.Command)
	GetOutput() BlockletOutput
	Options() RenderOptions
	EventManager() EventManager
}

type BlockletOutput interface {
	Print(data Data)
}
