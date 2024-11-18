package ui

import (
	"strconv"
	"strings"

	"github.com/adam-lavrik/go-imath/ix"
	Models "github.com/cyklan/photosynthesis/internal/models"
	"github.com/cyklan/photosynthesis/internal/ui/chars"
	"github.com/fatih/color"
)

func RenderGrid(game *Models.Game) string {
	var builder strings.Builder

	bgDefault := color.New(color.Reset).SprintFunc()
	bgRed := color.New(color.BgRed).SprintFunc()
	bgGreen := color.New(color.BgGreen).SprintFunc()
	backgroundColor := bgDefault

	for r := 0; r < game.Board.HexSize; r++ {
		actualR := r - Models.HexRadius
		qLength := Models.BoardHeight - ix.Abs(actualR)

		builder.WriteString(backgroundColor(strings.Repeat(" ", ix.Abs(actualR))))
		for q := 0; q < qLength; q++ {
			actualQ := ix.Max(-Models.HexRadius, -r) + q
			coord := Models.HexCoordinate{
				Q: actualQ,
				R: actualR,
			}

			cell := game.Board.Grid[coord]

			if cell.IsInShadow {
				backgroundColor = bgRed
			}

			if len(cell.CanPlant) > 0 {
				backgroundColor = bgGreen
			}

			builder.WriteString(backgroundColor(getPlayerColor(cell.Tree.Player)(chars.GetTreeChar(cell.Tree))))

			backgroundColor = bgDefault

			if q+1 != qLength {
				builder.WriteString(backgroundColor(" "))
			}
		}
		builder.WriteString("\n")
	}

	return builder.String()
}

func RenderGeneralData(game *Models.Game) string {
	var builder strings.Builder

	builder.WriteString("\n " + chars.SunChar + " " + chars.GetSunChars(game.SunState))
	builder.WriteString("\n Rounds remaining: " + strconv.Itoa(game.RemainingRounds))

	return builder.String()
}

func RenderPlayerData(game *Models.Game) string {
	var builder strings.Builder

	builder.WriteString("\n")

	for i, player := range game.Players {
		builder.WriteString(getPlayerColor(player.Id)(chars.SunChar + " Player " + strconv.Itoa(i) + ": " + strconv.Itoa(player.SunEnergy)))
		builder.WriteString("\n")
	}

	return builder.String()
}

func getPlayerColor(id int) func(a ...interface{}) string {
	switch id {
	case 0:
		return color.New(color.FgYellow).SprintFunc()
	case 1:
		return color.New(color.FgBlue).SprintFunc()
	case 2:
		return color.New(color.FgCyan).SprintFunc()
	case 3:
		return color.New(color.FgHiMagenta).SprintFunc()
	default:
		return color.New(color.FgWhite).SprintFunc()
	}
}
