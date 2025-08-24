package protocols

import (
	"os"
	"time"

	event "github.com/spetix/days2xmasleft/internal/data/event"
	render "github.com/spetix/days2xmasleft/internal/data/render"
	"github.com/spetix/days2xmasleft/internal/dateutil"
)

func ExampleRawOut_Print_ok() {
	out := NewRawOut(os.Stdout)
	desiredDate := time.Now().Add(11 * dateutil.Day).Format("01-02")
	daysToXmas := event.New("xmas", desiredDate, &render.RenderOptions{Label: "test", Format: "text", Unit: dateutil.Day}, time.Now)
	out.Print(daysToXmas)
	// Output:
	// 10
	// 10 to xmas
}

// func ExampleRawOut_Print_ok_withColor() {
// 	out := NewRawOut(os.Stdout)
// 	daysToXmas := time.Duration(10 * dateutil.Day)
// 	out.Print(daysToXmas, &RenderOptions{Label: "test", Format: "text", Unit: dateutil.Day, BackgroundColor: "#ff0000", ForegroundColor: "#00ff00"})
// 	// Output:
// 	// test10
// 	// test10
// 	// #00ff00
// 	// #ff0000

// }
