package day2

func GetResponseScore(input string) int {
	response := input[len(input)-1]
	value := Decode(response)
	return value
}

func GetResultScore(input string) int {
	return 1
}
