package day2

func Decode(character byte) int {
	value := 0
	switch character {
	case 'X', 'A':
		value = 1
	case 'Y', 'B':
		value = 2
	case 'Z', 'C':
		value = 3
	}

	return value
}

func GetWinRequirement(challenge int) int {
	requirement := challenge + 1
	// check for win
	if challenge == 3 {
		requirement = 1
	}
	return requirement
}
