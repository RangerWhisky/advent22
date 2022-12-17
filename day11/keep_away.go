package day11

import (
	"localhost/advent22/utils"
)

func SetupGame(filepath string) []Monkey {
	var monkeyList []Monkey

	input := utils.Read_file(filepath)
	for i := 1; i+5 <= len(input); i += 7 {
		monkeyDetails := input[i : i+5]
		monkeyList = append(monkeyList, InitMonkey(monkeyDetails))
	}

	return monkeyList
}
