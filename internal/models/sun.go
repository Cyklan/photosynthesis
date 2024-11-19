package models

type SunState int

const (
	TopRight SunState = iota
	Right
	BottomRight
	BottomLeft
	Left
	TopLeft
)

var directions = map[SunState]HexCoordinate{
	TopRight:    {Q: -1, R: 1},
	Right:       {Q: -1, R: 0},
	BottomRight: {Q: 0, R: -1},
	BottomLeft:  {Q: 1, R: -1},
	Left:        {Q: 1, R: 0},
	TopLeft:     {Q: 0, R: 1},
}

var SunStateCount = len(directions)

func (sunState SunState) getShadowCoords(origin HexCoordinate, treeHeight int) []HexCoordinate {
	shadowCoords := make([]HexCoordinate, 0, treeHeight)
	direction := directions[sunState]

	for i := 0; i < treeHeight; i++ {
		nextCoord := HexCoordinate{
			Q: origin.Q + ((i + 1) * direction.Q),
			R: origin.R + ((i + 1) * direction.R),
		}

		if nextCoord.GetDistanceFromCenter() > 3 {
			return shadowCoords
		}

		shadowCoords = append(shadowCoords, nextCoord)
	}

	return shadowCoords
}
