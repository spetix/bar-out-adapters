package protocols

import (
	"io"
	"os"

	"github.com/spetix/bar-out-adapters/pkg/barout/data"
)

type RawOut struct {
	baseOutput
}

func NewRawOut(device *os.File) *RawOut {
	return &RawOut{
		baseOutput: baseOutput{
			Device: device,
		},
	}
}

func (r *RawOut) Print(d data.Data) {

	writer := io.Writer(r.Device)
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
