package day10

const LineLength int = 40

func GetCrtOutput(cycle int, registerValue int) string {
	str := ""
	if IsPixelLit(cycle, registerValue) {
		str = "#"
	} else {
		str = "."
	}

	if getWidthFromCycle(cycle) == LineLength {
		str = str + "\n"
	}
	return str
}

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
