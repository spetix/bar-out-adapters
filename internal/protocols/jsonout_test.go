package protocols

import (
	"encoding/json"
	"os"
	"testing"
)

type fakeData struct{}

func (f *fakeData) Short() string           { return "name" }
func (f *fakeData) Long() string            { return "details" }
func (f *fakeData) Label() string           { return "name" }
func (f *fakeData) BackgroundColor() string { return "background" }
func (f *fakeData) ForegroundColor() string { return "foreground" }

func TestJsonOut_Print(t *testing.T) {
	tmp, err := os.CreateTemp("", "jsonout")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmp.Name())
	defer tmp.Close()

	j := NewJsonOut(tmp)

	m := &fakeData{}

	j.Print(m)
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
		t.Fatalf("text mismatch: got=%q", out["text"])
	}
	if out["tooltip"] != "details" {
		t.Fatalf("tooltip mismatch: got=%q", out["tooltip"])
	}
	if out["alt"] != "name" {
		t.Fatalf("alt mismatch: got=%q", out["alt"])
	}
}
