package day11

import "sort"

const NumRounds int = 20

func PartOne(filepath string) int {
	monkeyList := SetupGame(filepath, 3)

	for i := 0; i < NumRounds; i++ {
		monkeyList = TakeRound(monkeyList)
	}

	var itemPasses []int
	for i := 0; i < len(monkeyList); i++ {
		itemPasses = append(itemPasses, monkeyList[i].itemsInspected)
	}
	sort.Ints(itemPasses)

	return GetMonkeyBusiness(itemPasses)
}

func PartTwo(filepath string) int {
	monkeyList := SetupGame(filepath, 1)

	for i := 0; i < NumRounds; i++ {
		monkeyList = TakeRound(monkeyList)
	}

	var itemPasses []int
	for i := 0; i < len(monkeyList); i++ {
		itemPasses = append(itemPasses, monkeyList[i].itemsInspected)
	}
	sort.Ints(itemPasses)

	return GetMonkeyBusiness(itemPasses)
}

func GetMonkeyBusiness(values []int) int {
	last := len(values) - 1
	return values[last] * values[last-1]
}
