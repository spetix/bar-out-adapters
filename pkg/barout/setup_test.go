package barout

import "testing"

func TestBlockletProtocolSet(t *testing.T) {
	var proto BlockletProtocol
	if err := proto.Set("waybar"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if proto != ProtocolWaybar {
		t.Fatalf("expected protocol %q, got %q", ProtocolWaybar, proto)
	}
}

func TestNewReturnsOutputForWaybar(t *testing.T) {
	out := New("waybar")
	if out == nil {
		t.Fatal("expected non-nil output")
	}
}
