package day8

func PartOne(filepath string) int {
	forest := InitialiseForestFromFile(filepath)
	AssessVisibility(&forest)

	return GetVisibleTreeCount(&forest)
}
