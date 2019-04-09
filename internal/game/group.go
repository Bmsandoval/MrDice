package game

type Group struct {
	Players []Player
}

/**
Generator Function
 */
func CreateGroup(playerCount int) Group {
	group := Group{}

	// Generate playerCount players
	for p := 0; p < playerCount; p++{
		group.Players = append(group.Players, Player{
			p,
			Dice{},
			0,
			0,
		})
	}

	return group
}

