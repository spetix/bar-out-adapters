package tests

import (
	"github.com/spetix/bar-out-adapters/pkg/barout"
)

type dummydata struct {
	name       string
	details    string
	background string
	foreground string
}

func (d *dummydata) Short() string {
	return d.name
}
func (d *dummydata) Long() string {
	return d.details
}
func (d *dummydata) BackgroundColor() string {
	return d.background
}
func (d *dummydata) ForegroundColor() string {
	return d.foreground
}
func (d *dummydata) Label() string {
	return d.name
}

func ExampleBlockletOutput_ok() {

	barout.New("waybar").Print(&dummydata{"name", "details", "background", "foreground"})
	// 	Output:
	// {"text":"name","tooltip":"details","alt":"name","background-color":"background","foreground-color":"foreground"}

}
