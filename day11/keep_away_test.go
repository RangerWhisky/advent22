package day11

import "testing"

func TestMoneyGameSetup(t *testing.T) {
	monkeyList := SetupGame("./example.txt", 3)

	if len(monkeyList) != 4 {
		t.Error()
	}
}

func TestMoneyTurn(t *testing.T) {
	monkeyList := SetupGame("./example.txt", 3)

	monkeyList = TakeRound(monkeyList)

	//After round 1, the monkeys are holding items with these worry levels:
	// Monkey 0: 20, 23, 27, 26
	// Monkey 1: 2080, 25, 167, 207, 401, 1046
	// Monkey 2:
	// Monkey 3:

	expectedItems := []int{20, 23, 27, 26}

	for i, val := range monkeyList[0].items {
		if val != expectedItems[i] {
			t.Error()
		}
	}

	expectedItems = []int{2080, 25, 167, 207, 401, 1046}

	for i, val := range monkeyList[1].items {
		if val != expectedItems[i] {
			t.Error()
		}
	}

	if len(monkeyList[2].items) != 0 {
		t.Error()
	}
	if len(monkeyList[3].items) != 0 {
		t.Error()
	}
}

func TestNormalisationStep(t *testing.T) {
	monkeyList := SetupGame("./example.txt", 1)

	normalisedWorry := NormaliseWorryScore(&monkeyList[0], &monkeyList[1], 20)
	if normalisedWorry != 20 {
		t.Error()
	}

	normalisedWorry = NormaliseWorryScore(&monkeyList[0], &monkeyList[1], (2*19*23 + 1))
	if normalisedWorry != (19*23 + 1) {
		t.Error()
	}
}
