package protocols

import (
	"io"
	"os"

	"github.com/spetix/bar-out-adapters/pkg/barout/data"
)

type I3BlocksOut struct {
	baseOutput
}

func NewI3BlocksOut(device *os.File) *I3BlocksOut {
	return &I3BlocksOut{
		baseOutput: baseOutput{
			Device: device,
		},
	}
}

func (i *I3BlocksOut) Print(d data.Data) {

	writer := io.Writer(i.Device)
	//defer io.WriteCloser(r.Device).Close()

	// long and short format are identical

	// writer.Write([]byte(renderOptions.Label))
	writer.Write([]byte(d.Short()))
	writer.Write([]byte("\n"))
	writer.Write([]byte(d.Long()))
	writer.Write([]byte("\n"))

	if fgColor := d.ForegroundColor(); fgColor != "" {
		writer.Write([]byte(fgColor))
		writer.Write([]byte("\n"))
	}
	if bgColor := d.BackgroundColor(); bgColor != "" {
		writer.Write([]byte(bgColor))
		writer.Write([]byte("\n"))
	}
}
