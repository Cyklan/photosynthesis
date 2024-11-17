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

var stateChars = map[TreeState]string{
	Empty:   "⬡",
	Sapling: ".",
	Small:   "*",
	Medium:  "o",
	Large:   "O",
}

func (tree *Tree) GetChar() string {
	return stateChars[tree.TreeState]
}
