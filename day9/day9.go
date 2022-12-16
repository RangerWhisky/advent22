package day9

import (
	"localhost/advent22/utils"
)

type Coordinate struct {
	height, width int
}

func PartOne(filepath string) int {
	return ModelRopeMovement(filepath, 2)
}

func PartTwo(filepath string) int {
	return ModelRopeMovement(filepath, 10)
}

func ModelRopeMovement(filepath string, length int) int {
	maxHeight, maxWidth, startPosition := GetMapRequirements(filepath)
	visitMap := utils.InitialiseBoolMap(maxHeight, maxWidth)

	var rope []Coordinate
	for i := 0; i < length; i++ {
		rope = append(rope, startPosition)
	}

	MarkRopePosition(&visitMap, maxHeight, rope[0])

	for _, line := range utils.Read_file(filepath) {
		moveH, moveW := Decode(line)
		for incrementH := 1; incrementH <= moveH; incrementH++ {
			rope[0].height++
			rope = ResolveTail(rope, &visitMap, maxHeight)
		}
		for incrementH := -1; incrementH >= moveH; incrementH-- {
			rope[0].height--
			rope = ResolveTail(rope, &visitMap, maxHeight)
		}
		for incrementW := 1; incrementW <= moveW; incrementW++ {
			rope[0].width++
			rope = ResolveTail(rope, &visitMap, maxHeight)
		}

		for incrementW := -1; incrementW >= moveW; incrementW-- {
			rope[0].width--
			rope = ResolveTail(rope, &visitMap, maxHeight)
		}
	}
	return utils.GetMarkedSpaces(&visitMap)
}

func ResolveTail(rope []Coordinate, visitMap *utils.BoolMap, maxHeight int) []Coordinate {
	// for a rope greater than 2 points, this needs to be refactored into a loop
	for knot := 1; knot < len(rope); knot++ {
		positions := GetTailPositions(rope[knot], rope[knot-1])
		if len(positions) != 0 {
			rope[knot] = positions[0]
		}
		if knot == len(rope)-1 {
			// if the tail (len(rope) - 1) has moved, mark all the remaining positions
			for _, p := range positions {
				MarkRopePosition(visitMap, maxHeight, p)
			}
		}
	}
	return rope
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

func MarkRopePosition(visitMap *utils.BoolMap, maxHeight int, p Coordinate) {
	utils.MarkMap(visitMap, maxHeight-p.height-1, p.width)
}
