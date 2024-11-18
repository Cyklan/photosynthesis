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
		entry.Tree.TreeState = models.Medium
		entry.Tree.Player = 0
	}

	if entry, ok := game.Board.Grid[models.HexCoordinate{
		Q: 1,
		R: -1,
	}]; ok {
		entry.Tree.TreeState = models.Medium
		entry.Tree.Player = 1
	}

	game.SunState = models.TopLeft
	game.Update()

	fmt.Printf("\n%s\n", ui.RenderGrid(game))
	fmt.Println(ui.RenderGeneralData(game))
	fmt.Println(ui.RenderPlayerData(game))
}
