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
		t.Errorf("Response score is = %d, need 2", score)
	}
}

func TestGetScissorsResponseScore(t *testing.T) {
	testInput := "A Z"
	score := GetResponseScore(testInput)
	if score != 3 {
		t.Errorf("Response score is = %d, need 3", score)
	}
}

func TestGetLossResultScore(t *testing.T) {
	testInput := "A Z"
	score := GetResultScore(testInput)
	if score != 0 {
		t.Errorf("Response score is = %d, need 0", score)
	}
}

func TestGetDrawResultScore(t *testing.T) {
	testInput := "A X"
	score := GetResultScore(testInput)
	if score != 3 {
		t.Errorf("Response score is = %d, need 3", score)
	}
}

func TestGetWinResultScore(t *testing.T) {
	testInput := "A Y"
	score := GetResultScore(testInput)
	if score != 6 {
		t.Errorf("Response score is = %d, need 6", score)
	}
}
