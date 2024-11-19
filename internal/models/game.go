package models

import (
	"math/rand"
	"os"
)

const initialSunState SunState = TopRight

type Game struct {
	Board              Grid
	SunState           SunState
	RemainingRounds    int
	Players            []*Player
	ActivePlayer       *Player
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
		SunState:           initialSunState,
		RemainingRounds:    4,
		Players:            players,
		VictoryPointTokens: tokens,
		ActivePlayer:       players[0],
	}
}

func (game *Game) Init() {
	borderCells := game.Board.GetBorderCells()
	for i := 0; i < 2; i++ {
		for _, player := range game.Players {
			for {
				cellToPlaceOn := borderCells[rand.Intn(len(borderCells))]

				if cellToPlaceOn.Tree.TreeState != Empty {
					continue
				}

				cellToPlaceOn.Tree.Player = player.Id
				cellToPlaceOn.Tree.TreeState = Small

				break
			}
		}
	}
}

func (game *Game) Update() {
	game.Board.Update(game)
}

func (game *Game) NextTurn() {
	game.advancePlayer()
	game.advanceSunPosition()
	game.advanceTurn()

	game.checkForGameOver()
}

func (game *Game) advancePlayer() {
	nextPlayerIndex := (game.ActivePlayer.Id + 1) % len(game.Players)
	game.ActivePlayer = game.Players[nextPlayerIndex]
}

func (game *Game) advanceSunPosition() {
	if game.ActivePlayer.Id != game.Players[0].Id {
		return
	}

	game.SunState += 1
	if game.SunState == SunState(SunStateCount) {
		game.SunState = initialSunState
	}
}

func (game *Game) advanceTurn() {
	if game.SunState != initialSunState {
		return
	}

	game.RemainingRounds -= 1
}

func (game *Game) checkForGameOver() {
	if game.RemainingRounds == 0 {
		// todo: calculate and print out winner
		os.Exit(0)
	}
}
