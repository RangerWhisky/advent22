package day6

import (
	"bytes"
	file_input "localhost/advent22/utils"
)

const slidingWindow int = 4

func PartOne(filepath string) int {
	elfInput := file_input.Read_file(filepath)
	signal := []byte(elfInput[0])

	signalStart := slidingWindow

	for i := 0; i <= len(elfInput); i++ {
		signalStart = i + slidingWindow
		signalToCheck := signal[i:signalStart]
		if IsStartMarker(signalToCheck) {
			break
		}
	}
	return signalStart
}

func IsStartMarker(signal []byte) bool {
	isSignal := true

	for i := 1; i < len(signal); i++ {
		var subset []byte = signal[:i]
		var itemToCheck []byte = signal[i : i+1]
		if bytes.Contains(subset, itemToCheck) {
			isSignal = false
		}
	}
	return isSignal
}
