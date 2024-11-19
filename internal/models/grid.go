// https://www.redblobgames.com/grids/hexagons/#coordinates-axial

package models

import (
	"slices"

	"github.com/adam-lavrik/go-imath/ix"
)

type HexCoordinate struct {
	Q, R int
}

type GridCell struct {
	Tree       *Tree
	Leaves     int
	IsInShadow bool
	// List of player Ids
	CanPlant []int
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

func (coord HexCoordinate) GetNeighbours(distance int) []HexCoordinate {
	neighbours := []HexCoordinate{}

	// All hail the ai overlord
	for dq := -distance; dq <= distance; dq++ {
		for dr := max(-distance, -dq-distance); dr <= min(distance, -dq+distance); dr++ {
			neighbours = append(neighbours, HexCoordinate{
				Q: coord.Q + dq,
				R: coord.R + dr,
			})
		}
	}

	return neighbours
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

func (grid *Grid) GetBorderCells() []*GridCell {
  borderCells := []*GridCell{} 
  for coord, cell := range grid.Grid {
    if coord.GetDistanceFromCenter() == 3 {
      borderCells = append(borderCells, &cell)
    } 
  }

  return borderCells
}

func (grid *Grid) Update(game *Game) {
	grid.preUpdate(game)
	grid.update(game)
	grid.postUpdate(game)
}

func (grid *Grid) preUpdate(_ *Game) {
	for coord := range grid.Grid {
		grid.resetCell(coord)
	}
}

func (grid *Grid) update(game *Game) {
	for coord := range grid.Grid {
		grid.updateShadow(coord, game.SunState)
		grid.updateCanPlant(coord)
	}
}

func (grid *Grid) postUpdate(game *Game) {
	for coord := range grid.Grid {
		grid.updatePlayer(coord, game)
	}
}

func (grid *Grid) updatePlayer(coord HexCoordinate, game *Game) {
	cell := grid.Grid[coord]
	if cell.Tree.TreeState == Empty || cell.Tree.TreeState == Sapling || cell.IsInShadow {
		return
	}

	for _, player := range game.Players {
		if player.Id != cell.Tree.Player {
			continue
		}

		player.SunEnergy += int(cell.Tree.TreeState)
        if player.SunEnergy < 20 {
          player.SunEnergy = 20
        }

		return
	}
}

func (grid *Grid) resetCell(coord HexCoordinate) {
	if entry, ok := grid.Grid[coord]; ok {
		entry.IsInShadow = false
		entry.CanPlant = make([]int, 0, 4)

		grid.Grid[coord] = entry
	}
}

func (grid *Grid) updateShadow(coord HexCoordinate, sunState SunState) {
	shadowCaster := grid.Grid[coord]
	if shadowCaster.Tree.TreeState == Empty || shadowCaster.Tree.TreeState == Sapling {
		return
	}

	shadowLength := int(shadowCaster.Tree.TreeState)
	shadowCoords := sunState.getShadowCoords(coord, shadowLength)

	for _, shadowCoord := range shadowCoords {
		if entry, ok := grid.Grid[shadowCoord]; ok {
			treeHeight := int(entry.Tree.TreeState)

			if treeHeight <= shadowLength {
				entry.IsInShadow = true
			}

			grid.Grid[shadowCoord] = entry
		}
	}
}

func (grid *Grid) updateCanPlant(coord HexCoordinate) {
	cell := grid.Grid[coord]
	if cell.Tree.TreeState == Empty || cell.Tree.TreeState == Sapling || cell.IsInShadow {
		return
	}

	for _, neighbour := range coord.GetNeighbours(int(cell.Tree.TreeState)) {
		if entry, ok := grid.Grid[neighbour]; ok {
			if entry.IsInShadow || entry.Tree.TreeState != Empty {
				entry.CanPlant = make([]int, 0, 4)
			} else {
				entry.CanPlant = append(entry.CanPlant, cell.Tree.Player)
			}

			grid.Grid[neighbour] = entry
		}
	}
}

func (grid *Grid) GetPlantableCells(playerId int) []*GridCell {
  cells := []*GridCell{}
  for _ , cell := range grid.Grid {
    if slices.Contains(cell.CanPlant, playerId) {
      cells = append(cells, &cell) 
    }
  }

  return cells 
}

func (grid *Grid) GetPlayerTrees(playerId int) []*GridCell {
  cells := []*GridCell{}
  for _, cell := range grid.Grid {
    if cell.Tree.Player == playerId {
      cells = append(cells, &cell)
    }
  }

  return cells
} 

func (grid *Grid) GetScorableTrees(playerId int) []*GridCell {
  cells := []*GridCell{}
  for _, cell := range grid.GetPlayerTrees(playerId) {
     if cell.Tree.TreeState == Large {
       cells = append(cells, cell)
     }
  }
  
  return cells
}

