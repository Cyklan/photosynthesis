package models

type Player struct {
	Id int
	Tableau
	VictoryPointTokens []VictoryPointToken
	IsStartingPlayer   bool
	Bought             []Tree
}

func NewPlayer(id int, isStartingPlayer bool) Player {
	// seeds + small + medium + big
	bought := make([]Tree, 6+8+4+2)
	bought[0] = Tree{
		TreeState: Sapling,
		Player:    id,
	}
	bought[1] = Tree{
		TreeState: Sapling,
		Player:    id,
	}
	bought[2] = Tree{
		TreeState: Small,
		Player:    id,
	}
	bought[3] = Tree{
		TreeState: Small,
		Player:    id,
	}
	bought[4] = Tree{
		TreeState: Small,
		Player:    id,
	}
	bought[5] = Tree{
		TreeState: Small,
		Player:    id,
	}
	bought[6] = Tree{
		TreeState: Medium,
		Player:    id,
	}

	return Player{
		Id:                 id,
		Tableau:            NewTableau(),
		VictoryPointTokens: make([]VictoryPointToken, 0, 24),
		IsStartingPlayer:   isStartingPlayer,
		// seeds + small + medium + big
		Bought: bought,
	}
}