package day6

import (
	"testing"
)

func TestIsStartMarker(t *testing.T) {
	testData := []byte("mjqjpqmgbljsphdztnvjfqwrcgsmlb")

	subsetToCheck := testData[0:4]
	startMarker := IsStartMarker(subsetToCheck)

	if startMarker != false {
		t.Errorf("Found a start marker where there is none for %q", subsetToCheck)
	}

	subsetToCheck = testData[4:8]
	startMarker = IsStartMarker(subsetToCheck)
	if startMarker != true {
		t.Errorf("Failed to find a valid start marker for %q", subsetToCheck)
	}

}
