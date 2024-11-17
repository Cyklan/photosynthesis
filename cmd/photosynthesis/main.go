package main

import (
	"fmt"

	"github.com/cyklan/photosynthesis/internal/models"
	"github.com/cyklan/photosynthesis/internal/ui"
)

func main() {
	game := models.NewGame()

	if entry, ok := game.Board.Grid[models.HexCoordinate{
		Q: 0,
		R: 0,
	}]; ok {
		entry.Tree.TreeState = models.Small
	}

	if entry, ok := game.Board.Grid[models.HexCoordinate{
		Q: 3,
		R: 0,
	}]; ok {
		entry.Tree.TreeState = models.Sapling
	}

	if entry, ok := game.Board.Grid[models.HexCoordinate{
		Q: 0,
		R: 3,
	}]; ok {
		entry.Tree.TreeState = models.Medium
	}

	if entry, ok := game.Board.Grid[models.HexCoordinate{
		Q: 0,
		R: 2,
	}]; ok {
		entry.Tree.TreeState = models.Large
	}

	game.SunState = models.BottomRight
	game.Update()

	fmt.Printf("\n%s\n", ui.RenderGrid(game))
}
