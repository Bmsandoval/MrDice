package game

type Dice struct {
	dice []int
}

/**
Return the score of all dice
 */
func (d *Dice) Tally() int {
	score := 0
	for _, die := range d.dice {
		if die == 4 {
			// a roll of 4 counts as 0
			continue
		} else {
			score += die
		}
	}
	return score
}
