package day2

import (
	file_input "localhost/advent22/utils"
)

func PartOne(filepath string) int {
	runningTotal := 0
	elfInput := file_input.Read_file(filepath)
	for i := 0; i < len(elfInput); i++ {
		runningTotal += GetRoundScore(elfInput[i])
	}
	return runningTotal
}

func PartTwo(filepath string) int {
	runningTotal := 0
	elfInput := file_input.Read_file(filepath)
	for i := 0; i < len(elfInput); i++ {
		strategy := elfInput[i]
		action := GetActionForIntendedResult(elfInput[i])

		// Appending the new action to the end does the job purely because it becomes the last character
		// bit hacky, but if the behaviour doesn't change then the tests won't fail
		strategy = strategy + string(action)
		runningTotal += GetRoundScore(strategy)
	}
	return runningTotal
}

func GetRoundScore(input string) int {
	return GetResponseScore(input) + GetResultScore(input)
}

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

func GetActionForIntendedResult(input string) byte {
	// initialise response to meet the needs of a draw
	response := EncodeResponse(Decode(input[0]))
	challengeValue := Decode(input[0])
	// Check what we need and fix up our response
	switch input[len(input)-1] {
	case 'X':
		// need to lose, so get the loss requirement
		response = EncodeResponse(GetLossRequirement(challengeValue))
	case 'Z':
		response = EncodeResponse(GetWinRequirement(challengeValue))
	}
	return response
}
