package day9

import (
	"fmt"
	"localhost/advent22/utils"
)

type Coordinate struct {
	height, width int
}

func DayOne(filepath string) int {
	maxHeight, maxWidth := GetMapRequirements(filepath)
	visitMap := utils.InitialiseBoolMap(maxHeight, maxWidth)

	tailPosition := Coordinate{1, 1}
	headPosition := Coordinate{1, 1}
	utils.MarkMap(&visitMap, maxHeight-tailPosition.height, tailPosition.width-1)

	for _, line := range utils.Read_file(filepath) {
		moveH, moveW := Decode(line)
		fmt.Printf("\n%s\n", line)
		headPosition.height += moveH
		headPosition.width += moveW
		positions := GetTailPositions(tailPosition, headPosition)
		for _, p := range positions {
			utils.MarkMap(&visitMap, maxHeight-p.height, p.width-1)
			tailPosition = p
		}
		utils.PrintBoolMap(&visitMap)
	}
	return utils.GetMarkedSpaces(&visitMap)
}

func GetTailPositions(tail Coordinate, head Coordinate) []Coordinate {
	var positions []Coordinate
	diffH := head.height - tail.height
	diffW := head.width - tail.width

	// upwards movement case
	for h := 1; h < diffH; h++ {
		tailPosition := Coordinate{head.height - h, head.width}
		positions = append(positions, tailPosition)
	}

	// downwards movement case
	for h := -1; h > diffH; h-- {
		tailPosition := Coordinate{head.height - h, head.width}
		positions = append(positions, tailPosition)
	}

	// rightwards movement case
	for w := 1; w < diffW; w++ {
		tailPosition := Coordinate{head.height, head.width - w}
		positions = append(positions, tailPosition)
	}

	// rightwards movement case
	for w := -1; w > diffW; w-- {
		tailPosition := Coordinate{head.height, head.width - w}
		positions = append(positions, tailPosition)
	}

	return positions
}
