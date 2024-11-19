package models

const PlantSeedCost = 1
const GrowSeedCost = 1
const GrowSmallTreeCost = 2
const GrowMediumTreeCost = 3
const HarvestLargeTreeCost = 4
const MaxSunEnergy = 20

type Tableau struct {
	SunEnergy   int
	Seeds       []TableauSlot
	SmallTrees  []TableauSlot
	MediumTrees []TableauSlot
	LargeTrees  []TableauSlot
}

type TableauSlot struct {
	TreeState
	Cost        int
	IsAvailable bool
}

func NewTableau() Tableau {
	seeds := make([]TableauSlot, 4)
	seeds[0] = TableauSlot{
		TreeState:   Sapling,
		Cost:        1,
		IsAvailable: true,
	}
	seeds[1] = TableauSlot{
		TreeState:   Sapling,
		Cost:        1,
		IsAvailable: true,
	}
	seeds[2] = TableauSlot{
		TreeState:   Sapling,
		Cost:        2,
		IsAvailable: true,
	}
	seeds[3] = TableauSlot{
		TreeState:   Sapling,
		Cost:        3,
		IsAvailable: true,
	}

	smallTrees := make([]TableauSlot, 4)
	smallTrees[0] = TableauSlot{
		TreeState:   Small,
		Cost:        2,
		IsAvailable: true,
	}
	smallTrees[1] = TableauSlot{
		TreeState:   Small,
		Cost:        2,
		IsAvailable: true,
	}
	smallTrees[2] = TableauSlot{
		TreeState:   Small,
		Cost:        3,
		IsAvailable: true,
	}
	smallTrees[3] = TableauSlot{
		TreeState:   Small,
		Cost:        3,
		IsAvailable: true,
	}

	mediumTrees := make([]TableauSlot, 3)
	mediumTrees[0] = TableauSlot{
		TreeState:   Medium,
		Cost:        3,
		IsAvailable: true,
	}
	mediumTrees[1] = TableauSlot{
		TreeState:   Medium,
		Cost:        3,
		IsAvailable: true,
	}
	mediumTrees[2] = TableauSlot{
		TreeState:   Medium,
		Cost:        4,
		IsAvailable: true,
	}

	largeTrees := make([]TableauSlot, 2)
	largeTrees[0] = TableauSlot{
		TreeState:   Large,
		Cost:        4,
		IsAvailable: true,
	}
	largeTrees[1] = TableauSlot{
		TreeState:   Large,
		Cost:        5,
		IsAvailable: true,
	}

	return Tableau{
		SunEnergy:   0,
		Seeds:       seeds,
		SmallTrees:  smallTrees,
		MediumTrees: mediumTrees,
		LargeTrees:  largeTrees,
	}
}

func (tableau Tableau) GetVictoryPointsBySunEnergy() int {
	if tableau.SunEnergy < 3 {
		return 0
	}

	if tableau.SunEnergy < 6 {
		return 1
	}

	if tableau.SunEnergy < 9 {
		return 2
	}

	if tableau.SunEnergy < 12 {
		return 3
	}

	if tableau.SunEnergy < 15 {
		return 4
	}

	if tableau.SunEnergy < 18 {
		return 5
	}

	return 6
}
