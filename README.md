# uuid
A small, focused UUID v4 generator for Go.

This package provides a secure, RFC4122-compliant version 4 UUID generator with a tiny, easy-to-use API.

Highlights
- Uses crypto/rand for cryptographically secure randomness.
- Produces RFC4122 version 4 UUIDs (correct version and variant bits).
- Stateless, safe for concurrent use via the package-level Generate() function.
- Backwards-compatible convenience type UUID with New().Generate() retained for callers that expect it.

Quick start

Generate a UUID (recommended):

```go
package main

import (
	"fmt"
	"github.com/grimdork/uuid"
)

func main() {
	u, err := uuid.Generate()
	if err != nil {
		panic(err)
	}
	fmt.Println("UUID:", u)
}
```

Compatibility convenience (legacy API)

```go
u := uuid.New()
// Note: this convenience method returns a string and ignores rare crypto/rand errors.
id := u.Generate()
fmt.Println(id)
```

Behaviour and guarantees
- Format: 36-character lower-case hex string with hyphens, e.g. 123e4567-e89b-12d3-a456-426655440000
- Version: set to 4 (random UUID)
- Variant: RFC4122 variant (most-significant two bits = 10)
- Concurrency: package-level Generate() is safe for concurrent use. The legacy New().Generate() is a convenience wrapper and remains available.

Error handling
- Generate() returns an error only if crypto/rand fails to provide random bytes (extremely rare). Callers that need absolute guarantees should handle the error.

Notes for maintainers
- go.mod has been bumped to go 1.25.
- Tests validate formatting and RFC bits.

License
- See LICENSE for details.
