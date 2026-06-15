package protocols

import (
	"encoding/json"
	"log"
	"os"

	"github.com/spetix/bar-out-adapters/pkg/barout/data"
)

type JsonOut struct {
	baseOutput
}
type outputjson struct {
	Text            string `json:"text"`
	Tooltip         string `json:"tooltip"`
	Alt             string `json:"alt"`
	BackgroundColor string `json:"background-color"`
	ForegroundColor string `json:"foreground-color"`
}

func NewJsonOut(device *os.File) *JsonOut {
	return &JsonOut{
		baseOutput: baseOutput{
			Device: device,
		},
	}
}

func (j *JsonOut) Print(d data.Data) {
	newJson := outputjson{
		Text:            d.Short(),
		Tooltip:         d.Long(),
		Alt:             d.Label(),
		BackgroundColor: d.BackgroundColor(),
		ForegroundColor: d.ForegroundColor(),
	}
	b, err := json.Marshal(newJson)
	if err != nil {
		log.Print("json marshal error", err)
	}
	j.Device.Write(b)
}
