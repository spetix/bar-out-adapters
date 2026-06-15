package models

import (
	"testing"

	"github.com/spetix/bar-out-adapters/internal/protocols"
)

func TestBlockletProtocolSet(t *testing.T) {
	var proto protocols.BlockletProtocol
	if err := proto.Set("waybar"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if proto != protocols.ProtocolWaybar {
		t.Fatalf("expected protocol %q, got %q", protocols.ProtocolWaybar, proto)
	}
}
