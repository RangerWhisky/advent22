package day6

import (
	"bytes"
	file_input "localhost/advent22/utils"
)

const slidingWindow int = 4

func PartOne(filepath string) int {
	elfInput := file_input.Read_file(filepath)
	signal := []byte(elfInput[0])

	signalStart := 0

	for markerIndex := 0; markerIndex <= len(elfInput); markerIndex++ {
		signalToCheck := GetSlidingWindow(signal, markerIndex)
		if IsStartMarker(signalToCheck) {
			signalStart = markerIndex + slidingWindow
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

func GetSlidingWindow(signal []byte, start int) []byte {
	end := start + slidingWindow
	return signal[start:end]
}
