package game

import (
	"MrDice/pkg"
)

const MaxDieCount = 5
// Store data about each player
type Player struct {
	playerNumber int
	savedDice [MaxDieCount]int
	points int
}

/**
Return count of rollable dice
 */
func (p Player) rollableDice() int {
	// Count rollable dice
	dieCount := len(p.savedDice)
	rollableDie := MaxDieCount - dieCount
	return rollableDie
}

func (p Player) rollDice() []int {

	// Roll all unrolled dice for player
	// Return array of rolled dice
	return nil
}

func (p Player) saveDie(die int) error {
	if die < 1 || die > 6 {
		return pkg.Throw("Dice value out of range")
	}
	savableCount := p.rollableDice()
	if savableCount < 1 {
		return pkg.Throw("Out of rolls")
	}
	return nil
}

func (p Player) scoreRound() error {
	savableCount := p.rollableDice()
	if savableCount > 0 {
		return pkg.Throw("More dice to roll")
	}
	// Tally score and save
	return nil
}

