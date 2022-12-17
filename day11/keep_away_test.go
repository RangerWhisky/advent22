package day11

import "testing"

func TestMoneyTurn(t *testing.T) {
	monkeyList := SetupGame("./example.txt")

	if len(monkeyList) != 4 {
		t.Error()
	}
}
