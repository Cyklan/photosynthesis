package engine

import (
	"strings"

	"github.com/cyklan/photosynthesis/internal/game"
	"github.com/cyklan/photosynthesis/internal/math"
)

const hexRadius = 3
const hexSize = 7
const boardHeight = 7

type HexCoordinate struct {
	Q int
	R int
}

type GridCell struct {
	tree *game.Tree
}

type Grid struct {
	Grid map[HexCoordinate]GridCell
}

func NewGrid() *Grid {
	m := make(map[HexCoordinate]GridCell)

	for r := 0; r < hexSize; r++ {
		actualR := hexSize - hexRadius
		qLength := boardHeight - math.Abs(actualR)

		for q := 0; q < qLength; q++ {
			actualQ := math.Max(-hexRadius, -r)
			coord := HexCoordinate{
				Q: actualQ,
				R: actualR,
			}

			m[coord] = GridCell{}
		}
	}

	return &Grid{
		Grid: m,
	}
}

func (grid Grid) Render() string {
	var builder strings.Builder
	for r := -hexRadius; r <= hexRadius; r++ {
		absR := math.Abs(r)
		builder.WriteString(strings.Repeat(" ", absR))

		qLength := boardHeight - math.Abs(r)

		for q := 0; q < qLength; q++ {
			builder.WriteString("o")
			if q+1 != qLength {
				builder.WriteString(" ")
			}
		}

		builder.WriteString("\n")
	}

	return builder.String()
}

func (coord HexCoordinate) Neighbours() []HexCoordinate {
	return []HexCoordinate{
		{coord.Q + 1, coord.R},
		{coord.Q + 1, coord.R - 1},
		{coord.Q, coord.R - 1},
		{coord.Q - 1, coord.R},
		{coord.Q - 1, coord.R + 1},
		{coord.Q, coord.R + 1},
	}
}
