package protocols

import (
	"encoding/json"
	"log"
	"os"

	"github.com/spetix/bar-out-adapters/pkg/barout"
)

type WaybarOut struct {
	baseOutput
}
type outputjsonWaybar struct {
	Text            string `json:"text"`
	Tooltip         string `json:"tooltip"`
	Alt             string `json:"alt"`
	BackgroundColor string `json:"background-color"`
	ForegroundColor string `json:"foreground-color"`
}

func NewWaybarOut(device *os.File) barout.BlockletOutput {
	return &WaybarOut{
		baseOutput: baseOutput{
			Device: device,
		},
	}
}

func (j *WaybarOut) Print(d barout.Data) {
	newJson := outputjsonWaybar{
		Alt:             d.Label(),
		Text:            d.Short(), //dateutil.Format(unit, rerenderOptions.Unit),
		Tooltip:         d.Long(),  //fmt.Sprint(dateutil.Format(unit, rerenderOptions.Unit), " to xmas"),
		BackgroundColor: d.BackgroundColor(),
		ForegroundColor: d.ForegroundColor(),
	}
	b, err := json.Marshal(newJson)
	if err != nil {
		log.Print("json marshal error", err)
	}
	j.Device.Write(b)
}
