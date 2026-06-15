package protocols

import (
	"os"
	"strings"
	"testing"
)

type fakeDataRaw struct{}

func (f *fakeDataRaw) Short() string           { return "name" }
func (f *fakeDataRaw) Long() string            { return "details" }
func (f *fakeDataRaw) Label() string           { return "name" }
func (f *fakeDataRaw) BackgroundColor() string { return "background" }
func (f *fakeDataRaw) ForegroundColor() string { return "foreground" }

func TestRawOut_Print(t *testing.T) {
	tmp, err := os.CreateTemp("", "rawout")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmp.Name())
	defer tmp.Close()

	r := NewRawOut(tmp)
	d := &fakeDataRaw{}

	r.Print(d)
	tmp.Close()

	b, err := os.ReadFile(tmp.Name())
	if err != nil {
		t.Fatal(err)
	}
	s := string(b)
	lines := strings.Split(strings.TrimRight(s, "\n"), "\n")
	if len(lines) < 4 {
		t.Fatalf("unexpected lines output: %q", s)
	}
	if lines[0] != "name" {
		t.Fatalf("short mismatch: %q", lines[0])
	}
	if lines[1] != "details" {
		t.Fatalf("long mismatch: %q", lines[1])
	}
	if lines[2] != "foreground" {
		t.Fatalf("foreground mismatch: %q", lines[2])
	}
	if lines[3] != "background" {
		t.Fatalf("background mismatch: %q", lines[3])
	}
}
