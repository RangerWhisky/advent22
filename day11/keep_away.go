package day11

import (
	"localhost/advent22/utils"
)

func SetupGame(filepath string, worryDivisor int) []Monkey {
	var monkeyList []Monkey

	input := utils.Read_file(filepath)
	for i := 1; i+5 <= len(input); i += 7 {
		monkeyDetails := input[i : i+5]
		monkeyList = append(monkeyList, InitMonkey(monkeyDetails, worryDivisor))
	}

	return monkeyList
}

func TakeRound(monkeyList []Monkey) []Monkey {

	for i := 0; i < len(monkeyList); i++ {
		passList := TakeTurn(&monkeyList[i])
		for _, pass := range passList {
			// value := pass.val
			value := NormaliseWorryScore(monkeyList, pass.val)
			monkeyList[pass.target].items = append(monkeyList[pass.target].items, value)
		}
	}

	return monkeyList
}

func NormaliseWorryScore(monkeys []Monkey, score int) int {
	lcm := 1
	newScore := score
	for i := 0; i < len(monkeys); i++ {
		lcm = lcm * monkeys[i].divisor
	}
	if score > lcm {
		newScore = (score % lcm)
	}

	return newScore
}
