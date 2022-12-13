package day9

import "localhost/advent22/utils"

type Coordinate struct {
	height, width int
}

func DayOne(filepath string) int {
	visitMap := utils.InitialiseBoolMap(GetMapRequirements(filepath))

	tailPosition := Coordinate{1, 1}
	headPosition := Coordinate{1, 1}

	for _, line := range utils.Read_file(filepath) {
		moveH, moveW := Decode(line)
		headPosition.height += moveH
		headPosition.width += moveW
		positions := GetTailPositions(tailPosition, headPosition)
		for _, p := range positions {
			utils.MarkMap(&visitMap, p.height, p.width)
		}
	}

	return 0
}

func GetTailPositions(tail Coordinate, head Coordinate) []Coordinate {
	var positions []Coordinate
	positions = append(positions, Coordinate{1, 1})
	return positions
}
