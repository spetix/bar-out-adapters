package protocols

import (
	"encoding/json"
	"log"
	"os"

	"github.com/spetix/bar-out-adapters/pkg/barout"
)

type JsonOut struct {
	baseOutput
}
type outputjson struct {
	Short           string `json:"short"`
	Long            string `json:"long"`
	Label           string `json:"label"`
	BackgroundColor string `json:"background-color"`
	ForegroundColor string `json:"foreground-color"`
}

func NewJsonOut(device *os.File) barout.BlockletOutput {
	return &JsonOut{
		baseOutput: baseOutput{
			Device: device,
		},
	}
}

func (j *JsonOut) Print(d barout.Data) {
	newJson := outputjson{
		Label:           d.Label(),
		Short:           d.Short(),
		Long:            d.Long(),
		BackgroundColor: d.BackgroundColor(),
		ForegroundColor: d.ForegroundColor(),
	}
	b, err := json.Marshal(newJson)
	if err != nil {
		log.Print("json marshal error", err)
	}
	j.Device.Write(b)
}
