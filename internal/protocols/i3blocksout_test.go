package protocols

import (
	"os"
	"strings"
	"testing"
)

type fakeDataI3 struct{}

func (f *fakeDataI3) Short() string           { return "name" }
func (f *fakeDataI3) Long() string            { return "details" }
func (f *fakeDataI3) Label() string           { return "name" }
func (f *fakeDataI3) BackgroundColor() string { return "background" }
func (f *fakeDataI3) ForegroundColor() string { return "foreground" }

func TestI3BlocksOut_Print(t *testing.T) {
	tmp, err := os.CreateTemp("", "i3out")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmp.Name())
	defer tmp.Close()

	i := NewI3BlocksOut(tmp)
	d := &fakeDataI3{}

	i.Print(d)
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
	if lines[0] != "name" || lines[1] != "details" {
		t.Fatalf("short/long mismatch: %v", lines[:2])
	}
}
