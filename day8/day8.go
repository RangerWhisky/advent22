package day8

func PartOne(filepath string) int {
	forest := InitialiseForestFromFile(filepath)
	AssessVisibility(&forest)

	return GetVisibleTreeCount(&forest)
}

func PartTwo(filepath string) int {
	forest := InitialiseForestFromFile(filepath)

	h, w := GetForestDimensions(&forest)
	bestScenicScore := 1

	for i := 1; i < h; i++ {
		for j := 1; j < w; j++ {
			scenicScore := GetScenicScore(&forest, i, j)
			if scenicScore > bestScenicScore {
				bestScenicScore = scenicScore
			}
		}
	}

	return bestScenicScore
}
