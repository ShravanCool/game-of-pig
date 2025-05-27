package src

import "math/rand"

func rollDice() int {
	return rand.Intn(6) + 1
}
