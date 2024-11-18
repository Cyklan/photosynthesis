package models

import (
	"math/rand"
)

type Game struct {
	Board              Grid
	SunState           SunState
	RemainingRounds    int
	Players            []*Player
	VictoryPointTokens map[int][]VictoryPointToken
}

func NewGame() *Game {
	players := make([]*Player, 4)
	for i := range players {
		players[i] = NewPlayer(i, i == 0)
	}
	tokens := make(map[int][]VictoryPointToken)

	tokens[1] = []VictoryPointToken{
		NewVictoryPointToken(1, 12),
		NewVictoryPointToken(1, 12),
		NewVictoryPointToken(1, 12),
		NewVictoryPointToken(1, 12),
		NewVictoryPointToken(1, 13),
		NewVictoryPointToken(1, 13),
		NewVictoryPointToken(1, 13),
		NewVictoryPointToken(1, 14),
		NewVictoryPointToken(1, 14),
	}

	tokens[2] = []VictoryPointToken{
		NewVictoryPointToken(1, 13),
		NewVictoryPointToken(1, 13),
		NewVictoryPointToken(1, 14),
		NewVictoryPointToken(1, 14),
		NewVictoryPointToken(1, 16),
		NewVictoryPointToken(1, 16),
		NewVictoryPointToken(1, 17),
	}

	tokens[3] = []VictoryPointToken{
		NewVictoryPointToken(1, 17),
		NewVictoryPointToken(1, 17),
		NewVictoryPointToken(1, 18),
		NewVictoryPointToken(1, 18),
		NewVictoryPointToken(1, 19),
	}

	tokens[4] = []VictoryPointToken{
		NewVictoryPointToken(1, 20),
		NewVictoryPointToken(1, 21),
		NewVictoryPointToken(1, 22),
	}

	return &Game{
		Board:              *NewGrid(),
		SunState:           TopRight,
		RemainingRounds:    4,
		Players:            players,
		VictoryPointTokens: tokens,
	}
}

func (game *Game) Init() {
  borderCells := game.Board.GetBorderCells()
  for i := 0; i < 2; i++ {
    for player := 1; player <= 4; player++ {
      for {
        cellToPlaceOn := borderCells[rand.Intn(len(borderCells))] 

        if cellToPlaceOn.Tree.TreeState != Empty {
          continue
        }

        cellToPlaceOn.Tree.Player = player
        cellToPlaceOn.Tree.TreeState = Small 

        break
      }
    }
  }
}

func (game *Game) Update() {
	game.Board.Update(game)
}
