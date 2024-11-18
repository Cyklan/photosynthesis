package models

type TreeState int

type Tree struct {
	Player int
	TreeState
}

const (
	Empty TreeState = iota - 1
	Sapling
	Small
	Medium
	Large
)

func NewTree(player int) Tree {
	return Tree{
		Player:    player,
		TreeState: Sapling,
	}
}