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
		headPositions := GetHeadPositions(rope[0], line)
		for _, value := range headPositions {
			rope[0] = value
			rope = ResolveRopeForStep(rope, &visitMap, maxHeight)
		}
	}
	return utils.GetMarkedSpaces(&visitMap)
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

func GetHeadPositions(head Coordinate, line string) []Coordinate {
	var positions []Coordinate
	diffH, diffW := Decode(line)

	// upwards movement case
	for h := 1; h <= diffH; h++ {
		head := Coordinate{head.height + h, head.width}
		positions = append(positions, head)
	}

	// downwards movement case
	for h := -1; h >= diffH; h-- {
		head := Coordinate{head.height + h, head.width}
		positions = append(positions, head)
	}

	// rightwards movement case
	for w := 1; w <= diffW; w++ {
		head := Coordinate{head.height, head.width + w}
		positions = append(positions, head)
	}

	// rightwards movement case
	for w := -1; w >= diffW; w-- {
		head := Coordinate{head.height, head.width + w}
		positions = append(positions, head)
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
		newTail.height = head.height
		newTail.width += stepW
	}
	return newTail
}

func getStepToClose(diff int) int {
	step := diff

	if step >= 2 {
		step = step - 1
	} else if step <= -2 {
		step = step + 1
	} else {
		step = 0
	}
	return step
}

func MarkRopePosition(visitMap *utils.BoolMap, maxHeight int, p Coordinate) {
	utils.MarkMap(visitMap, maxHeight-p.height-1, p.width)
}
