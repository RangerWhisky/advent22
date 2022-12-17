package day10

const LineLength int = 40

func GetCrtOutput(pixel int, registerValue int) string {
	str := ""
	if IsPixelLit(pixel, registerValue) {
		str = "#"
	} else {
		str = "."
	}

	if getWidthFromPixel(pixel) == LineLength-1 {
		str = str + "\n"
	}
	return str
}

func IsPixelLit(pixel int, registerValue int) bool {
	width := getWidthFromPixel(pixel)
	if width >= registerValue-1 && width <= registerValue+1 {
		return true
	}
	return false
}

func getWidthFromPixel(pixel int) int {
	return (pixel % LineLength)
}
