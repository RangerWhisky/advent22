package day6

import "testing"

func TestIsStartMarker(t *testing.T) {
	testData := []byte("mjqjpqmgbljsphdztnvjfqwrcgsmlb")
	startMarker := IsStartMarker(testData[0:4])

	if startMarker != false {
		t.Errorf("Found a start marker where there is none")
	}
}
