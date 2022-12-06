package day6

import (
	"bytes"
)

func PartOne(filepath string) int {
	return 0
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
