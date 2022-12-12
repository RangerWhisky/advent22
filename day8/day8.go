package day8

func PartOne(filepath string) int {
	forest := InitialiseForestFromFile("./simplified_example.txt")
	AssessVisibility(&forest)

	return GetVisibleTreeCount(&forest)
}
