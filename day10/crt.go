package day10

const LineLength int = 40

func IsPixelLit(cycle int, registerValue int) bool {
	width := getWidthFromCycle(cycle)
	if width >= registerValue-1 && width <= registerValue+1 {
		return true
	}
	return false
}

func getWidthFromCycle(cycle int) int {
	return ((cycle - 1) % LineLength) + 1
}
