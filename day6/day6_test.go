package day6

import (
	"bytes"
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

func TestCornerCaseStartMarker(t *testing.T) {
	testData := []byte("mjqjpqmgj")

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

func TestGetSlidingWindow(t *testing.T) {

	testData := []byte("mjqjpqmgj")

	subsetToCheck := GetSlidingWindow(testData, 0)
	expectedSlice := testData[0:4]

	if bytes.Compare(expectedSlice, subsetToCheck) != 0 {
		t.Errorf("Requested subsets not equal.  Got %q expected %q", subsetToCheck, expectedSlice)
	}
}

func TestPartOneSolution(t *testing.T) {
	solution := PartOne("./simplified_example.txt")
	if solution != 10 {
		t.Errorf("Start of packet marker complete after %d characters but should be %d", solution, 10)
	}
}
