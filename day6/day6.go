package day6

import "bytes"

func IsStartMarker(signal []byte) bool {
	isSignal := true

	for i := 1; i < len(signal); i++ {
		var subset []byte = signal[0:i]
		var itemToCheck []byte = signal[i:i]
		if bytes.Contains(subset, itemToCheck) {
			isSignal = false
		}
	}
	return isSignal
}
