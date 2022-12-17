package day11

import (
	"strings"
	"testing"
)

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

func TestGetFalseDestination(t *testing.T) {
	input := "    If false: throw to monkey 3"

	monkey := GetThrowDestination(input)

	if monkey != 3 {
		t.Error()
	}
}

func TestOperationInitiation(t *testing.T) {
	input := "	Operation: new = old + 6"

	operation := GetOperation(input)

	example := PerformOperation(operation, 7)

	if example != 13 {
		t.Error()
	}
}

func TestPower(t *testing.T) {
	input := "	Operation: new = old * old"

	operation := GetOperation(input)

	example := PerformOperation(operation, 7)

	if example != 49 {
		t.Error()
	}
}

func TestSubtract(t *testing.T) {
	input := "	Operation: new = old - 4"

	operation := GetOperation(input)

	example := PerformOperation(operation, 7)

	if example != 3 {
		t.Error()
	}
}

func TestMultiply(t *testing.T) {
	input := "	Operation: new = old * 3"

	operation := GetOperation(input)

	example := PerformOperation(operation, 7)

	if example != 21 {
		t.Error()
	}
}

func TestInitMonkey(t *testing.T) {
	input := `Starting items: 79, 98
	Operation: new = old * 19
	Test: divisible by 23
	  If true: throw to monkey 2
	  If false: throw to monkey 3
  `
	monkey := InitMonkey(strings.Split(input, "\n"))

	Inspect(&monkey, 0)

	if monkey.items[0] != (79 * 19 / 3) {
		t.Error()
	}
}

func TestThrowTrue(t *testing.T) {
	input := `Starting items: 79, 98
	Operation: new = old * 19
	Test: divisible by 23
	  If true: throw to monkey 2
	  If false: throw to monkey 3
  `
	monkey := InitMonkey(strings.Split(input, "\n"))

	Inspect(&monkey, 0)

	target := GetThrowTarget(&monkey, 0)

	if target != 3 {
		t.Error()
	}
}
