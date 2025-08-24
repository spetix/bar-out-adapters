package protocols

// func ExampleJsonOut_Print_ok() {
// 	out := NewJsonOut(os.Stdout)
// 	desiredDate := time.Now().Add(11 * dateutil.Day).Format("01-02")
// 	e := event.New("xmas", desiredDate, &render.RenderOptions{Label: "test", Format: "text", Unit: dateutil.Day}, time.Now)
// 	out.Print(e)
// 	// Output:
// 	// {"short":"10","long":"10 to xmas","label":"test","background-color":"","foreground-color":""}
// }

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
