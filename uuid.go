package uuid

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// Generate returns a RFC4122-compliant version 4 UUID as a string.
// It is stateless, uses crypto/rand, and is safe for concurrent use.
func Generate() (string, error) {
	var id [16]byte
	_, err := rand.Read(id[:])
	if err != nil {
		return "", fmt.Errorf("uuid: crypto/rand failed: %w", err)
	}
	// Set RFC4122 variant and version bits.
	id[6] = (id[6] & 0x0f) | 0x40 // Version 4
	id[8] = (id[8] & 0x3f) | 0x80 // Variant is 10xxxxxx

	// Efficient hex encoding
	s := make([]byte, 36)
	hex.Encode(s[0:8], id[0:4])
	s[8] = '-'
	hex.Encode(s[9:13], id[4:6])
	s[13] = '-'
	hex.Encode(s[14:18], id[6:8])
	s[18] = '-'
	hex.Encode(s[19:23], id[8:10])
	s[23] = '-'
	hex.Encode(s[24:36], id[10:16])
	return string(s), nil
}

// Keep a backward-compatible UUID type with Generate method that delegates to the package Generate.
// This avoids breaking callers that used uuid.New().Generate().

type UUID struct{}

func New() *UUID { return &UUID{} }

func (u *UUID) Generate() string {
	s, _ := Generate()
	return s
}
