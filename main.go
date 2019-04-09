package main

import (
	"MrDice/internal"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	internal.PlayGame()
}
