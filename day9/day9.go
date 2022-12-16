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
			rope = ResolveRopeForStep(rope, &visitMap, maxHeight)
		}
		for incrementH := -1; incrementH >= moveH; incrementH-- {
			rope[0].height--
			rope = ResolveRopeForStep(rope, &visitMap, maxHeight)
		}
		for incrementW := 1; incrementW <= moveW; incrementW++ {
			rope[0].width++
			rope = ResolveRopeForStep(rope, &visitMap, maxHeight)
		}
		for incrementW := -1; incrementW >= moveW; incrementW-- {
			rope[0].width--
			rope = ResolveRopeForStep(rope, &visitMap, maxHeight)
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

func ResolveRopeForStep(rope []Coordinate, visitMap *utils.BoolMap, maxHeight int) []Coordinate {
	for knot := 1; knot < len(rope); knot++ {
		position := GetSnapPosition(rope[knot], rope[knot-1])
		if rope[knot] == position {
			break
		}
		rope[knot] = position
		if knot == len(rope)-1 {
			MarkRopePosition(visitMap, maxHeight, rope[knot])
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

func GetSnapPosition(tail Coordinate, head Coordinate) Coordinate {
	diffH := head.height - tail.height
	diffW := head.width - tail.width

	newTail := tail

	stepH := getStepToClose(diffH)
	stepW := getStepToClose(diffW)

	if stepH != 0 && stepW != 0 {
		newTail.height += stepH
		newTail.width += stepW
	} else if stepH != 0 {
		newTail.height += stepH
		newTail.width = head.width
	} else if stepW != 0 {
		newTail.width += stepH
		newTail.height = head.height
	}
	return newTail
}

func getStepToClose(diff int) int {
	step := diff
	switch diff {
	case 2, -2:
		// little move
		step = diff / 2
	case 1, -1:
		// only move if width stretches
	default:
		// already in line
	}
	return step
}

func MarkRopePosition(visitMap *utils.BoolMap, maxHeight int, p Coordinate) {
	utils.MarkMap(visitMap, maxHeight-p.height-1, p.width)
}
