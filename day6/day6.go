package day6

import (
	"bytes"
	"fmt"
)

func IsStartMarker(signal []byte) bool {
	isSignal := true

	for i := 1; i < len(signal); i++ {
		var subset []byte = signal[:i]
		var itemToCheck []byte = signal[i : i+1]
		if bytes.Contains(subset, itemToCheck) {
			fmt.Printf("Is %q in %q", itemToCheck, subset)
			isSignal = false
		}
	}
	return isSignal
}
