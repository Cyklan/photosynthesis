package models

import (
	"fmt"

	"github.com/adam-lavrik/go-imath/ix"
)

type HexCoordinate struct {
	Q, R int
}

type GridCell struct {
	Tree       *Tree
	Leaves     int
	IsInShadow bool
}

type Grid struct {
	HexRadius   int
	HexSize     int
	BoardHeight int
	Grid        map[HexCoordinate]GridCell
}

const HexRadius = 3
const HexSize = 7
const BoardHeight = 7

func NewGrid() *Grid {
	m := make(map[HexCoordinate]GridCell)

	for r := 0; r < HexSize; r++ {
		actualR := r - HexRadius
		qLength := BoardHeight - ix.Abs(actualR)

		for q := 0; q < qLength; q++ {
			actualQ := ix.Max(-HexRadius, -r) + q
			coord := HexCoordinate{
				Q: actualQ,
				R: actualR,
			}

			m[coord] = GridCell{
				Leaves: coord.GetLeaves(),
				Tree: &Tree{
					TreeState: Empty,
					Player:    -1,
				},
				IsInShadow: false,
			}
		}
	}

	return &Grid{
		HexRadius:   HexRadius,
		HexSize:     HexSize,
		BoardHeight: BoardHeight,
		Grid:        m,
	}
}

func (coord HexCoordinate) GetNeighbours() []HexCoordinate {
	return []HexCoordinate{
		{coord.Q + 1, coord.R},
		{coord.Q + 1, coord.R - 1},
		{coord.Q, coord.R - 1},
		{coord.Q - 1, coord.R},
		{coord.Q - 1, coord.R + 1},
		{coord.Q, coord.R + 1},
	}
}

func (coord HexCoordinate) GetLeaves() int {
	return 4 - coord.GetDistanceFromCenter()
}

func (coord HexCoordinate) GetDistanceFromCenter() int {
	center := HexCoordinate{
		Q: 0,
		R: 0,
	}

	distanceQ := center.Q - coord.Q
	distanceR := center.R - coord.R

	return (ix.Abs(distanceQ) + ix.Abs(distanceR) + ix.Abs(distanceQ+distanceR)) / 2
}

func (grid *Grid) Update(game *Game) {
	for coord := range grid.Grid {
		if entry, ok := grid.Grid[coord]; ok {
			entry.IsInShadow = false

			grid.Grid[coord] = entry
		}
	}
	
	for coord := range grid.Grid {
		grid.updateShadow(coord, game.SunState)
	}
}

func (grid *Grid) updateShadow(coord HexCoordinate, sunState SunState) {
	shadowCaster := grid.Grid[coord]
	if shadowCaster.Tree.TreeState == Empty || shadowCaster.Tree.TreeState == Sapling {
		return
	}

	shadowLength := int(shadowCaster.Tree.TreeState)
	shadowCoords := sunState.getShadowCoords(coord, shadowLength)

	fmt.Println(shadowCoords)

	for _, shadowCoord := range shadowCoords {
		if entry, ok := grid.Grid[shadowCoord]; ok {
			treeHeight := int(entry.Tree.TreeState)

			if treeHeight < shadowLength {
				entry.IsInShadow = true
			}

			grid.Grid[shadowCoord] = entry
		}
	}
}
