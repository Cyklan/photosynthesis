package main

import (
	"fmt"

	"github.com/cyklan/photosynthesis/internal/models"
	"github.com/cyklan/photosynthesis/internal/ui"
)

func main() {
	game := models.NewGame()

	game.Init()

	game.SunState = models.TopLeft
	game.Update()

	fmt.Printf("\n%s\n", ui.RenderGrid(game))
	fmt.Println(ui.RenderGeneralData(game))
	fmt.Println(ui.RenderPlayerData(game))
}
