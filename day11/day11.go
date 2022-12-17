package day11

import (
	"fmt"
	"sort"
)

const NumRounds int = 20
const PartTwoRounds int = 10000

func PartOne(filepath string) int {
	return RunGame(filepath, 3, NumRounds)
}

func PartTwo(filepath string) int {
	return RunGame(filepath, 1, PartTwoRounds)
}

func RunGame(filepath string, calmFactor int, rounds int) int {
	monkeyList := SetupGame(filepath, calmFactor)

	for i := 0; i < rounds; i++ {
		monkeyList = TakeRound(monkeyList)
	}

	var itemPasses []int
	for i := 0; i < len(monkeyList); i++ {
		fmt.Printf("Monkey %d has %d passes\n", i, monkeyList[i].itemsInspected)
		itemPasses = append(itemPasses, monkeyList[i].itemsInspected)
	}
	sort.Ints(itemPasses)

	return GetMonkeyBusiness(itemPasses)
}

func GetMonkeyBusiness(values []int) int {
	last := len(values) - 1
	return values[last] * values[last-1]
}
