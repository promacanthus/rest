# rest

## Usages

```go
package main

import (
	"github.com/promacanthus/rest/example"
	"github.com/promacanthus/rest/rest"
)

func main() {
	cli := example.NewClient(&rest.Config{
		BaseURLs: map[string]string{
			"Name": "http://localhost:8080",
		},
	})

	// use cli ...
}
```