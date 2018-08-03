# Golang Game Loop

[![GoDoc](https://godoc.org/github.com/kutase/go-gameloop?status.svg)](https://godoc.org/github.com/kutase/go-gameloop)

> :video_game: :arrows_counterclockwise: Golang Game Loop implementation

## Install

```bash
go get github.com/parth/go-gameloop
```

## TODO

Add explanation for what the 10 argument means. I had to check, other people will likely have to as well

## Example

```go
package main

import (
	"github.com/parth/go-gameloop"
	"log"
)

func main() {
	callsPerSecond := 10
	
	gl := gameLoop.New(callsPerSecond, func(delta float64) {
		log.Println("tick:", delta)
	})

	gl.Start()

	// Stop Game Loop:
	// gl.Stop()

	// Don't stop main goroutine
	for {}
}
```
