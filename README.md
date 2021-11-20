# uuid
A simple UUID generator package for Go.

## How
Simplest usage:
```go
package main

import (
	"fmt"
	"github.com/grimdork/uuid"
)

func main() {
	u := uuid.New.Generate()
	fmt.Printf("UUID: %s\n", u)
}
```
