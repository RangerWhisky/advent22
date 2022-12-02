package day2

func GetResponseScore(input string) int {
	response := input[len(input)-1]
	value := 0
	switch response {
	case 'X':
		value = 1
	case 'Y':
		value = 2
	case 'Z':
		value = 3
	}
	return value
}

func GetResultScore(input string) int {
	return 1
}
