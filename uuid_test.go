package uuid

import (
	"strings"
	"testing"
)

func TestGenerateFormatAndBits(t *testing.T) {
	s, err := Generate()
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}
	if len(s) != 36 {
		t.Fatalf("unexpected length: %d", len(s))
	}
	// hyphen positions
	if s[8] != '-' || s[13] != '-' || s[18] != '-' || s[23] != '-' {
		t.Fatalf("hyphens not in expected positions: %s", s)
	}
	// version character at position 14 (zero-based) corresponds to id[6] high nibble
	if s[14] != '4' {
		t.Fatalf("expected version 4 at position 14, got %c in %s", s[14], s)
	}
	// variant: first hex char of segment starting at 19 (s[19]) has high bits 8..b (1000..1011)
	c := s[19]
	if !strings.Contains("89ab", strings.ToLower(string(c))) {
		t.Fatalf("expected variant in [89ab], got %c in %s", c, s)
	}
}

// Backward-compatibility test removed: legacy New().Generate() API has been removed.
func TestGenerateSimple(t *testing.T) {
	s, err := Generate()
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}
	if len(s) != 36 {
		t.Fatalf("unexpected length: %d", len(s))
	}
}
