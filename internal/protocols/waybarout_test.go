package protocols

import (
	"io"
	"os"
	"testing"
)

type testData struct {
	name       string
	details    string
	background string
	foreground string
}

func (d *testData) Short() string {
	return d.name
}

func (d *testData) Long() string {
	return d.details
}

func (d *testData) BackgroundColor() string {
	return d.background
}

func (d *testData) ForegroundColor() string {
	return d.foreground
}

func (d *testData) Label() string {
	return d.name
}

func TestWaybarOutPrint(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()

	out := NewWaybarOut(w)
	out.Print(&testData{"name", "details", "background", "foreground"})
	w.Close()

	got, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}

	want := `{"text":"name","tooltip":"details","alt":"name","background-color":"background","foreground-color":"foreground"}`
	if string(got) != want {
		t.Fatalf("unexpected output: got %q want %q", string(got), want)
	}
}
