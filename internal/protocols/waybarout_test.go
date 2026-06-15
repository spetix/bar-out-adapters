package protocols

import (
	"encoding/json"
	"os"
	"testing"
)

type fakeDataWaybar struct{}

func (f *fakeDataWaybar) Short() string           { return "name" }
func (f *fakeDataWaybar) Long() string            { return "details" }
func (f *fakeDataWaybar) Label() string           { return "name" }
func (f *fakeDataWaybar) BackgroundColor() string { return "background" }
func (f *fakeDataWaybar) ForegroundColor() string { return "foreground" }

func TestWaybarOut_Print(t *testing.T) {
	tmp, err := os.CreateTemp("", "waybarout")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmp.Name())
	defer tmp.Close()

	w := NewWaybarOut(tmp)
	d := &fakeDataWaybar{}

	w.Print(d)
	tmp.Close()

	b, err := os.ReadFile(tmp.Name())
	if err != nil {
		t.Fatal(err)
	}

	var out map[string]string
	if err := json.Unmarshal(b, &out); err != nil {
		t.Fatalf("invalid json: %v, raw=%s", err, string(b))
	}
	if out["text"] != "name" {
		t.Fatalf("text mismatch: %q", out["text"])
	}
	if out["tooltip"] != "details" {
		t.Fatalf("tooltip mismatch: %q", out["tooltip"])
	}
	if out["alt"] != "name" {
		t.Fatalf("alt mismatch: %q", out["alt"])
	}
}
