package day9

import (
	"localhost/advent22/utils"
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
		heightDiff = magnitude
	case "D":
		heightDiff = -magnitude
	}

	return heightDiff, widthDiff
}

func GetMapRequirements(filepath string) (int, int) {
	instructions := utils.Read_file(filepath)

	// assume instructions start in the bottom left so expect a right and/or up first
	currentHeight, currentWidth, maxHeight, maxWidth := 1, 1, 1, 1

	for _, line := range instructions {
		height, width := Decode(line)

		currentHeight += height
		currentWidth += width

		if currentHeight > maxHeight {
			maxHeight = currentHeight
		}

		if currentWidth > maxWidth {
			maxWidth = currentWidth
		}
	}

	return maxHeight, maxWidth
}
