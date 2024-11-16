package main

import (
	"fmt"

	"github.com/cyklan/photosynthesis/internal/engine"
)

func main() {
	grid := engine.NewGrid()
	fmt.Printf("\n%s\n", grid.Render())
}
