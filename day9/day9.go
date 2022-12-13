package day9

import "localhost/advent22/utils"

func DayOne(filepath string) int {
	// visitMap := utils.InitialiseBoolMap(GetMapRequirements(filepath))

	tailHeight, tailWidth := 1, 1
	headHeight, headWidth := 1, 1

	for _, line := range utils.Read_file(filepath) {
		moveH, moveW := Decode(line)
		headHeight += moveH
		headWidth += moveW
		GetTailPositions(tailHeight, tailWidth, headHeight, headWidth)

	}

	return 0
}

func GetTailPositions(tH int, tW int, hH int, hW int) []int {
	var positions []int
	positions = append(positions, 1, 1)
	return positions
}
