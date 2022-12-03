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
	incorrectItem := GetCompartmentDuplicate(backpack)
	return GetItemPriority(incorrectItem)
}

func GetCompartmentDuplicate(backpack string) byte {
	comparmentSize := getCompartmentSize(backpack)
	dupeLetters := getDuplicatedLetters(string(backpack[:comparmentSize]), string(backpack[comparmentSize:]))
	return dupeLetters[0]
}

func getCompartmentSize(backpack string) int {
	return len(backpack) / 2
}

func getDuplicatedLetters(left string, right string) []byte {
	var dupeLetters []byte

	for i := 0; i < len(left); i++ {
		matchCharacter := byte(left[i])
		for j := 0; j < len(right); j++ {
			if matchCharacter == byte(right[j]) {
				dupeLetters = append(dupeLetters, matchCharacter)
			}
		}
	}
	return dupeLetters
}

func GetBadgeLetter(backpacks []string) byte {
	firstSet := getDuplicatedLetters(backpacks[0], backpacks[1])
	finalSet := getDuplicatedLetters(string(firstSet), backpacks[2])

	return finalSet[0]
}
