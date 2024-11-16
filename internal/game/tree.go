package game

type TreeState int

type Tree struct {
	Player int
	TreeState
}

const (
	Sapling TreeState = iota
	Small
	Medium
	Large
)

func newTree() Tree {
	return Tree{
		Player:    1,
		TreeState: Sapling,
	}
}
