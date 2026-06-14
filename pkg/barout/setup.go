package barout

import (
	"os"

	"github.com/spetix/bar-out-adapters/internal/protocols"
	"github.com/spf13/cobra"
)

// RenderOptions contains display configuration for the OTP blocklet output.
type RenderOptions struct {
	Label           string
	Format          string
	ForegroundColor string
	BackgroundColor string
}


type Formatter interface {
	Render(s ...string) string
	Validate(f string) bool
}

type SetupBlocklet struct {
	options   RenderOptions
	protocol  BlockletProtocol
	formatter Formatter
}

func NewSetupBlocklet(f Formatter) *SetupBlocklet {
	return &SetupBlocklet{
		options: RenderOptions{
			Label:           "🎄",
			Format:          "",
			ForegroundColor: "#ffffff",
			BackgroundColor: "#000000",
		},
		formatter: f,
		protocol:  ProtocolRaw,
	}
}

func (s *SetupBlocklet) Setup(cmd *cobra.Command) {
	if proto := os.Getenv("protocol"); proto != "" {
		s.protocol.Set(proto)
	}
	cmd.PersistentFlags().Var(&s.protocol, "protocol", "output protocol (raw, i3blocks, waybar, json)")
	s.options.Format = ""
	if txtFmt := os.Getenv("format"); txtFmt != "" {
		s.options.Format = txtFmt
	}
	cmd.PersistentFlags().StringVar(&s.options.Format, "format", s.options.Format, "format")

	if envLbl := os.Getenv("label"); envLbl != "" {
		s.options.Label = envLbl
	}
	cmd.PersistentFlags().StringVar(&s.options.Label, "label", s.options.Label, "label")

	if envClr := os.Getenv("color"); envClr != "" {
		s.options.ForegroundColor = envClr
	}
	cmd.PersistentFlags().StringVar(&s.options.ForegroundColor, "color", s.options.ForegroundColor, "color")

	if envBgClr := os.Getenv("bgcolor"); envBgClr != "" {
		s.options.BackgroundColor = envBgClr
	}
	cmd.PersistentFlags().StringVar(&s.options.BackgroundColor, "bgcolor", s.options.BackgroundColor, "background color")
}

// New creates a new BlockletOutput based on the provided protocol string.
// It supports "json", "raw", and "waybar" protocols. If an unsupported protocol is provided, it defaults to "raw".
func (s *SetupBlocklet) GetOutput() BlockletOutput {

	switch s.protocol {
	case ProtocolRaw, ProtocolI3Blocks:
		return protocols.NewRawOut(os.Stdout)
	case ProtocolWaybar:
		return protocols.NewWaybarOut(os.Stdout)
	case ProtocolJson:
		return protocols.NewJsonOut(os.Stdout)
	default:
		return protocols.NewRawOut(os.Stdout)
	}
}

func New(protocol string) BlockletOutput {
	switch protocol {
	case string(ProtocolI3Blocks), string(ProtocolRaw):
		return protocols.NewRawOut(os.Stdout)
	case string(ProtocolWaybar):
		return protocols.NewWaybarOut(os.Stdout)
	case string(ProtocolJson):
		return protocols.NewJsonOut(os.Stdout)
	default:
		return protocols.NewRawOut(os.Stdout)
	}
}
