package internal

import (
	"MrDice/internal/game"
	"fmt"
	"math/rand"
	"sort"
)

const PlayerCount = 4
const RoundCount = 4
const DiceCount = 5

func PlayGame() {
	// Setup Game
	group := game.CreateGroup(PlayerCount)

	// Select a random starting player for round 1's playerOne

	startingPlayer := rand.Intn(3)
	// Track the hiscore to find the winners
	for r := 1; r <= RoundCount; r++ {
		// Each player gets to be playerOne once during a game
		playerID := (startingPlayer+r)%PlayerCount
		winningPlayers := PlaySingleRound(playerID, group)
		fmt.Printf("With %v points, the winner(s) of round %v are %v\n",
			group.Players[winningPlayers[0]].RoundPoints,
			r,
			winningPlayers)
	}

	var bestScore = 100
	var winners []int
	for playerID,player := range group.Players {
		if player.Points < bestScore {
			// track the winner
			winners = winners[:0]
			bestScore = player.Points
			winners = append(winners, playerID)
		} else if player.Points == bestScore {
			// handle ties
			winners = append(winners, playerID)
		}
	}
	fmt.Println("Game Over!")
	fmt.Printf("With %v points, the winner(s) are %v\n",
		bestScore,
		winners)

	// Lowest score of each round wins, but the final winner is who matters
	// Return winning player and score
	return
}
/*
Plays one of four rounds
Returns ids of winning players
 */
func PlaySingleRound(playerOne int, group game.Group) []int {
	maxTurns := DiceCount
	var bestScore = 100
	var winners []int
	for p := 0; p < PlayerCount; p++ {
		// Map rolling starting player to an array index
		playerID := (playerOne+p)%PlayerCount

		player := group.Players[playerID]
		player.RoundPoints = 0
		for t:=0; t < maxTurns; t++ {
			// Stop if player has saved all their dice
			if player.RollableDice() == 0 {
				break
			}
			// Otherwise roll available dice
			dice := RollDice(player)
			selectedDice := selectDice(dice)
			player.SaveDice(selectedDice)
			fmt.Printf("player %v rolled %v and kept %v\n", playerID, dice, selectedDice)
		}
		player.RoundPoints = player.SavedDice.Tally()
		player.ResetDice()
		fmt.Println(group.Players[playerID])
		player.UpdateTotal()
		group.Players[playerID] = player
		fmt.Println(group.Players[playerID])
		if player.RoundPoints < bestScore {
			// track the winner
			winners = winners[:0]
			bestScore = player.RoundPoints
			winners = append(winners, playerID)
		} else if player.RoundPoints == bestScore {
			// handle ties
			winners = append(winners, playerID)
		}
	}
	return winners
}

type byPoints []int

func (p byPoints) Len() int {
	return len(p)
}

func (p byPoints) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

/**
Sort dice by potential score
 */
func (p byPoints) Less(i, j int) bool {
	// 4 is actually less than anything else
	if p[i] == 4 {
		return true
	}
	return p[i] < p[j]
}

func selectDice(dice []int) []int {
	sort.Sort(byPoints(dice))
	var selected []int
	var lowestDie int
	for _, die := range dice {
		// Let's prefer potential scores of 0 or 1
		if die == 1 || die == 4 {
			selected = append(selected, die)
		} else {
			lowestDie = die
			break
		}
	}
	// Select at least one die each roll
	if len(selected) == 0 {
		selected = append(selected, lowestDie)
	}
	return selected
}

/**
For given player, roll all remaining unrolled dice
 */
func RollDice(p game.Player) []int {
	// Count remaining rollable dice
	rollableDie := p.RollableDice()
	if rollableDie == 0 {
		return nil
	}

	// Roll all unrolled dice for player
	var dice []int
	for i := 1; i<= rollableDie; i++ {
		// Roll d6
		die := rand.Intn(6 - 1) + 1
		dice = append(dice, die)
	}

	// Return array of rolled dice
	return dice
}

