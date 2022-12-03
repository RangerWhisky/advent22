package day3

import file_input "localhost/advent22/utils"

func PartOne(filepath string) int {
	runningTotal := 0
	elfInput := file_input.Read_file(filepath)

	for i := 0; i < len(elfInput); i++ {
		runningTotal += GetIncorrectItemPriority(elfInput[i])
	}

	return runningTotal
}

func GetIncorrectItemPriority(backpack string) int {
	incorrectItem := GetDuplicatedLetters(backpack)
	return GetItemPriority(incorrectItem)
}

func GetDuplicatedLetters(backpack string) byte {
	dupeLetter := byte('a')
	comparmentSize := getCompartmentSize(backpack)

	for i := 0; i < getCompartmentSize(backpack); i++ {
		matchCharacter := byte(backpack[i])
		for j := comparmentSize; j < len(backpack); j++ {
			if matchCharacter == byte(backpack[j]) {
				dupeLetter = matchCharacter
			}
		}
	}

	return dupeLetter
}

func getCompartmentSize(backpack string) int {
	return len(backpack) / 2
}

func GetBadgeLetter(backpacks []string) byte {
	return 'a'
}
