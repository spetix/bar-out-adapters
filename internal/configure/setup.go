package configure

import (
	"os"

	"github.com/spetix/bar-out-adapters/internal/events"
	"github.com/spetix/bar-out-adapters/internal/eventstrategy"
	"github.com/spetix/bar-out-adapters/internal/protocols"
	"github.com/spetix/bar-out-adapters/internal/render"
	"github.com/spetix/bar-out-adapters/pkg/barout/models"
	"github.com/spf13/cobra"
)

type SetupBlockletImpl struct {
	models.SetupBlocklet
	options   *render.RenderOptionsImpl
	protocol  protocols.BlockletProtocol
	formatter models.Formatter
	eventMgr  *events.EventManagerImpl
}

func NewSetupBlockletImpl(f models.Formatter) *SetupBlockletImpl {
	return &SetupBlockletImpl{
		options:   render.NewRenderOptionsImpl(),
		formatter: f,
		protocol:  protocols.ProtocolRaw,
		eventMgr:  events.NewEventManagerImpl(),
	}
}

func (s *SetupBlockletImpl) Setup(cmd *cobra.Command) {

	if proto := os.Getenv("protocol"); proto != "" {
		s.protocol.Set(proto)
	}
	cmd.PersistentFlags().Var(&s.protocol, "protocol", "output protocol (raw, i3blocks, waybar, json)")
	if txtFmt := os.Getenv("format"); txtFmt != "" {
		s.options.SetFormat(&txtFmt)
	}
	s.options.SetFormat(cmd.PersistentFlags().String("format", s.options.Format(), "format"))

	if envLbl := os.Getenv("label"); envLbl != "" {
		s.options.SetLabel(&envLbl)
	}
	s.options.SetLabel(cmd.PersistentFlags().String("label", s.options.Label(), "label"))

	if envClr := os.Getenv("color"); envClr != "" {
		s.options.SetForegroundColor(&envClr)
	}
	s.options.SetForegroundColor(cmd.PersistentFlags().String("color", s.options.ForegroundColor(), "color"))

	if envBgClr := os.Getenv("bgcolor"); envBgClr != "" {
		s.options.SetBackgroundColor(&envBgClr)
	}
	s.options.SetBackgroundColor(cmd.PersistentFlags().String("bgcolor", s.options.BackgroundColor(), "background color"))
}

// New creates a new BlockletOutput based on the provided protocol string.
// It supports "json", "raw", and "waybar" protocols. If an unsupported protocol is provided, it defaults to "raw".
func (s *SetupBlockletImpl) GetOutput() models.BlockletOutput {

	switch s.protocol {
	case protocols.ProtocolRaw, protocols.ProtocolI3Blocks:
		return protocols.NewRawOut(os.Stdout)
	case protocols.ProtocolWaybar:
		return protocols.NewWaybarOut(os.Stdout)
	case protocols.ProtocolJson:
		return protocols.NewJsonOut(os.Stdout)
	default:
		return protocols.NewRawOut(os.Stdout)
	}
}

func (s *SetupBlockletImpl) EventManager() models.EventManager {
	switch s.protocol {
	case protocols.ProtocolRaw, protocols.ProtocolI3Blocks:
		eventstrategy.NewEnvVarStrategy(s.eventMgr)
	case protocols.ProtocolWaybar:
		eventstrategy.NewNopStrategy(s.eventMgr)
	default:
		eventstrategy.NewNopStrategy(s.eventMgr)
	}
	return s.eventMgr
}

func (s *SetupBlockletImpl) Options() models.RenderOptions {
	return s.options
}
