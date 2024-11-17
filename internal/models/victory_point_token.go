package models

type VictoryPointToken struct {
	Leaves int
	Value  int
}

func NewVictoryPointToken(leaves int, value int) VictoryPointToken {
	return VictoryPointToken{
		Leaves: leaves,
		Value:  value,
	}
}