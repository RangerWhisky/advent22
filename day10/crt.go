package day10

func IsPixelLit(cycle int, registerValue int) bool {
	if cycle >= registerValue-1 && cycle <= registerValue+1 {
		return true
	}
	return false
}
