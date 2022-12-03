package day3

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
