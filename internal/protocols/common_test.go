package protocols

import "testing"

func TestBlockletProtocol_StringAndType(t *testing.T) {
	var p BlockletProtocol = ProtocolRaw
	if p.String() != "raw" {
		t.Fatalf("String() expected 'raw', got %q", p.String())
	}
	if p.Type() != "BlockletProtocol" {
		t.Fatalf("Type() expected 'BlockletProtocol', got %q", p.Type())
	}
}

func TestBlockletProtocol_Set(t *testing.T) {
	var p BlockletProtocol
	if err := p.Set("json"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if p != ProtocolJson {
		t.Fatalf("expected ProtocolJson, got %v", p)
	}

	if err := p.Set("i3blocks"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if p != ProtocolI3Blocks {
		t.Fatalf("expected ProtocolI3Blocks, got %v", p)
	}

	// invalid protocol should set to raw and return error
	if err := p.Set("unknown"); err == nil {
		t.Fatalf("expected error for unknown protocol")
	}
	if p != ProtocolRaw {
		t.Fatalf("expected ProtocolRaw on invalid set, got %v", p)
	}

	// cover explicit raw
	if err := p.Set("raw"); err != nil {
		t.Fatalf("unexpected error for raw: %v", err)
	}
	if p != ProtocolRaw {
		t.Fatalf("expected ProtocolRaw for raw set, got %v", p)
	}
}
