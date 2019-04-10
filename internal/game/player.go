package game

import (
	"math/rand"
	"time"
)

const DieCount = 5
// Store data about each player
type Player struct {
	playerNumber int
	SavedDice Dice
	RoundPoints int
	Points int
}

/**
Return count of rollable dice
 */
func (p Player) RollableDice() int {
	// Count rollable dice
	dieCount := len(p.SavedDice.dice)
	rollableDie := DieCount - dieCount
	return rollableDie
}

/**
For given player, roll all remaining unrolled dice
 */
func (p *Player) RollDice() []int {
	rand.Seed(time.Now().UnixNano())
	// Count remaining rollable dice
	rollableDie := p.RollableDice()
	if rollableDie == 0 {
		return nil
	}

	// Roll all unrolled dice for player
	var dice []int
	for i := 1; i<= rollableDie; i++ {
		// Roll d6
		die := rand.Intn(6) + 1
		dice = append(dice, die)
	}

	// Return array of rolled dice
	return dice
}

/**
Save one die for player P
 */
func (p *Player) SaveDice(dice []int) {
	for _, die:= range dice {
		if die < 1 || die > 6 {
			panic("Dice value out of range")
		}
		p.SavedDice.dice = append(p.SavedDice.dice, die)
	}
	return
}

/**
Save one die for player P
 */
func (p *Player) ResetDice() {
	p.SavedDice.dice = p.SavedDice.dice[:0]
	return
}

func (p *Player) SaveRoundPoints(points int) {
	p.RoundPoints = points
	return
}

/**
Save players score, tallying their dice for this round
 */
func (p *Player) UpdateTotal() {
	p.Points += p.RoundPoints
	return
}

