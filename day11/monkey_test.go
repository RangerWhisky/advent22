package day11

import "testing"

func TestInitialiseStartingItems(t *testing.T) {
	input := "	Starting items: 79, 98"
	itemList := GetStartingItems(input)

	if len(itemList) != 2 {
		t.Error()
	}

	if itemList[0] != 79 {
		t.Error()
	}
	if itemList[1] != 98 {
		t.Error()
	}
}

func TestInitialiseTestDivisor(t *testing.T) {
	input := "  Test: divisible by 19"
	divisor := GetTestDivisor(input)

	if divisor != 19 {
		t.Error()
	}
}

func TestGetTrueDestination(t *testing.T) {
	input := "    If true: throw to monkey 2"

	monkey := GetThrowDestination(input)

	if monkey != 2 {
		t.Error()
	}
}
