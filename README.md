# gamp [![GoDoc](https://godoc.org/github.com/olebedev/go-gamp?status.svg)](https://godoc.org/github.com/olebedev/go-gamp)

> [Google Analytics Measurement Protocol](https://developers.google.com/analytics/devguides/collection/protocol/v1/reference) in Golang

Almost full API implementation, except dynamic parameters(due to swagger 2.0 is not supported it yet) and batch mode.

### Example

```golang
package main

import (
	"log"
	"context"

	"github.com/AlekSi/pointer"
	gamp "github.com/olebedev/go-gamp"
	"github.com/olebedev/go-gamp/client/gampops"
)

func main() {
	client := gamp.New(context.Background(), "UA-XXXXXXXX-X")
	err := client.Collect(
		gampops.NewCollectParams().
			WithCid(pointer.ToString("42")).
			WithT("event").
			WithEc(pointer.ToString("Category")).
			WithEa(pointer.ToString("Action")),
	)
	if err != nil {
		log.Fatal(err)
	}
}
```
