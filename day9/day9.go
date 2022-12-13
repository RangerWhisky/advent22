package day9

import "localhost/advent22/utils"

type Coordinate struct {
	x, y int
}

func DayOne(filepath string) int {
	// visitMap := utils.InitialiseBoolMap(GetMapRequirements(filepath))

	tailPosition := Coordinate{1, 1}
	headPosition := Coordinate{1, 1}

	for _, line := range utils.Read_file(filepath) {
		moveH, moveW := Decode(line)
		headPosition.x += moveH
		headPosition.y += moveW
		GetTailPositions(tailPosition, headPosition)

	}

	return 0
}

func GetTailPositions(tail Coordinate, head Coordinate) []Coordinate {
	var positions []Coordinate
	positions = append(positions, Coordinate{1, 1})
	return positions
}
