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

	yellow := color.New(color.FgYellow).SprintFunc()

	bgDefault := color.New(color.Reset).SprintFunc()
	bgRed := color.New(color.BgRed).SprintFunc()
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

			builder.WriteString(backgroundColor(chars.GetTreeChar(cell.Tree)))

			backgroundColor = bgDefault

			if q+1 != qLength {
				builder.WriteString(backgroundColor(" "))
			}
		}
		builder.WriteString("\n")
	}

	builder.WriteString(yellow("\n " + chars.SunChar + " " + chars.GetSunChars(game.SunState)))
	builder.WriteString("\n Rounds remaining: " + strconv.Itoa(game.RemainingRounds))

	return builder.String()
}
