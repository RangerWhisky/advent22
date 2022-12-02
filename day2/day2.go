package day2

func GetResponseScore(input string) int {
	response := input[len(input)-1]
	value := Decode(response)
	return value
}

func GetResultScore(input string) int {
	// default to a loss
	score := 0

	challenge, response := GetChallengeAndResponse(input)

	// check for draw
	if challenge == response {
		score = 3
	}

	if response == GetWinRequirement(challenge) {
		score = 6
	}

	return score
}

func GetChallengeAndResponse(input string) (int, int) {
	challenge := Decode(input[0])
	response := Decode(input[len(input)-1])

	return challenge, response
}
