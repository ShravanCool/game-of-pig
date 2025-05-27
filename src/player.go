package src

type Player struct {
	Name             string
	HoldScore, Score int
}

func (p Player) getDecision(winningScore, turnScore int) string {
	if p.Score+turnScore >= winningScore {
		return "hold"
	} else if turnScore >= p.HoldScore {
		return "hold"
	} else {
		return "roll"
	}
}

func (p *Player) playTurn(winningScore int) int {
	turnScore := 0
	for p.getDecision(winningScore, turnScore) == "roll" {
		diceVal := rollDice()
		if diceVal == 1 {
			return 0
		}
		turnScore += diceVal
	}

	return turnScore
}
