package render

import "testing"

func TestRenderOptions_GettersSetters(t *testing.T) {
	r := NewRenderOptionsImpl()
	if r.Label() != "🎄" {
		t.Fatalf("default label mismatch: %q", r.Label())
	}
	if r.Format() != "" {
		t.Fatalf("default format mismatch: %q", r.Format())
	}
	if r.ForegroundColor() != "#ffffff" {
		t.Fatalf("default foreground mismatch: %q", r.ForegroundColor())
	}
	if r.BackgroundColor() != "#000000" {
		t.Fatalf("default background mismatch: %q", r.BackgroundColor())
	}

	lbl := "x"
	fmt := "txt"
	fg := "#111111"
	bg := "#222222"
	r.SetLabel(&lbl)
	r.SetFormat(&fmt)
	r.SetForegroundColor(&fg)
	r.SetBackgroundColor(&bg)

	if r.Label() != lbl {
		t.Fatalf("label after set mismatch: %q", r.Label())
	}
	if r.Format() != fmt {
		t.Fatalf("format after set mismatch: %q", r.Format())
	}
	if r.ForegroundColor() != fg {
		t.Fatalf("fg after set mismatch: %q", r.ForegroundColor())
	}
	if r.BackgroundColor() != bg {
		t.Fatalf("bg after set mismatch: %q", r.BackgroundColor())
	}
}
