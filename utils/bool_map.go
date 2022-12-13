package utils

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

}
