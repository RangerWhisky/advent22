package utils

import "fmt"

type BoolMap struct {
	dataPoints [][]bool
	countTrue  int
}

func InitialiseBoolMap(height int, width int) BoolMap {
	boolMap := new(BoolMap)
	for h := 0; h < height; h++ {
		var line []bool
		for w := 0; w < width; w++ {
			line = append(line, false)
		}
		boolMap.dataPoints = append(boolMap.dataPoints, line)
	}
	boolMap.countTrue = 0

	return *boolMap
}

func MarkMap(boolmap *BoolMap, height int, width int) {
	if !boolmap.dataPoints[height][width] {
		boolmap.dataPoints[height][width] = true
		boolmap.countTrue++
	}
}

func GetMarkedSpaces(boolmap *BoolMap) int {
	return boolmap.countTrue
}

func PrintBoolMap(boolmap *BoolMap) {
	height := len(boolmap.dataPoints)
	width := len(boolmap.dataPoints[0])
	for line := 0; line < height; line++ {
		for i := 0; i < width; i++ {
			if boolmap.dataPoints[line][i] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}
