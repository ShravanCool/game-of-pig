package src

import "testing"

var winningScore int = 100

func TestPlayerDecision(t *testing.T) {
	tests :=
		[]struct {
			name, expected string
			p              Player
			turnScore      int
		}{
			{
				name:     "Should return 'roll' when turn score is less than hold score",
				expected: "roll",
				p: Player{
					HoldScore: 20,
				},
				turnScore: 3,
			},
			{
				name:     "Should return 'hold' when turn score is greater than or equal to hold score",
				expected: "hold",
				p: Player{
					HoldScore: 20,
				},
				turnScore: 23,
			},
			{
				name:     "Should return 'hold' when turn score and total score is greater than or equal to winning score",
				expected: "hold",
				p: Player{
					HoldScore: 20,
					Score:     99,
				},
				turnScore: 3,
			},
		}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if decision := tt.p.getDecision(winningScore, tt.turnScore); decision != tt.expected {
				t.Errorf("player.getDecision() = %v, expected: %v", decision, tt.expected)
			}
		})
	}
}
