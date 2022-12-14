package day9

import (
	"fmt"
	"localhost/advent22/utils"
)

type Coordinate struct {
	height, width int
}

func PartOne(filepath string) int {
	maxHeight, maxWidth, startPosition := GetMapRequirements(filepath)
	visitMap := utils.InitialiseBoolMap(maxHeight, maxWidth)

	tailPosition := startPosition
	headPosition := startPosition
	utils.MarkMap(&visitMap, maxHeight-tailPosition.height-1, tailPosition.width)

	for _, line := range utils.Read_file(filepath) {
		moveH, moveW := Decode(line)
		fmt.Printf("\n%s\n", line)
		headPosition.height += moveH
		headPosition.width += moveW
		positions := GetTailPositions(tailPosition, headPosition)
		for _, p := range positions {
			utils.MarkMap(&visitMap, maxHeight-p.height-1, p.width)
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
