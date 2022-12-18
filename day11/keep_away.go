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
			value := pass.val
			// value := NormaliseWorryScore(&monkeyList[i], &monkeyList[pass.target], pass.val)
			monkeyList[pass.target].items = append(monkeyList[pass.target].items, value)
		}
	}

	return monkeyList
}

func NormaliseWorryScore(source *Monkey, dest *Monkey, score int) int {
	newScore := score % (source.divisor * dest.divisor)

	return newScore
}
