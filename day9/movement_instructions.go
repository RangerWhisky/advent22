package day9

import (
	"strconv"
	"strings"
)

func Decode(instruction string) (int, int) {

	heightDiff := 0
	widthDiff := 0

	stringParts := strings.Split(instruction, " ")
	magnitude, _ := strconv.Atoi(stringParts[1])

	switch stringParts[0] {
	case "R":
		widthDiff = magnitude
	case "L":
		widthDiff = -magnitude
	case "U":
		heightDiff = -magnitude
	}

	return heightDiff, widthDiff
}
