package src

import "testing"

func TestRollDice(t *testing.T) {
	expected := func(res int) bool { return res >= 1 && res <= 6 }

	res := rollDice()
	if !expected(res) {
		t.Errorf("rollDice()=%v, which isn't between 1 and 6", res)
	}
}
