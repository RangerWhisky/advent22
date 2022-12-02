package day2

import "testing"

func TestGetRockResponseScore(t *testing.T) {
	testInput := "A X"
	score := GetResponseScore(testInput)
	if score != 1 {
		t.Errorf("Response score is = %d, need 1", score)
	}
}

func TestGetPaperResponseScore(t *testing.T) {
	testInput := "A Y"
	score := GetResponseScore(testInput)
	if score != 2 {
		t.Errorf("Response score is = %d, need 1", score)
	}
}

func TestGetScissorsResponseScore(t *testing.T) {
	testInput := "A Z"
	score := GetResponseScore(testInput)
	if score != 3 {
		t.Errorf("Response score is = %d, need 1", score)
	}
}
