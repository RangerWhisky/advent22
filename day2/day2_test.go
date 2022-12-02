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

func TestSimpleGetRoundScore(t *testing.T) {
	testInput := "A Y"
	score := GetRoundScore(testInput)
	if score != 8 {
		t.Errorf("Response score is = %d, need 8", score)
	}
}

func TestGetActionForIntendedResult(t *testing.T) {
	testInput := "A Y"
	action := GetActionForIntendedResult(testInput)
	if action != 'X' {
		t.Errorf("Asked for a draw with Rock (X), but got %q", action)
	}
}

func TestEasySolution(t *testing.T) {
	solution := Easy("./simplified_example.txt")
	if solution != 15 {
		t.Errorf("Sample solution Score calculated as %d but should be %d", solution, 15)
	}
}

func TestBonusSolution(t *testing.T) {
	solution := Bonus("./simplified_example.txt")
	if solution != 12 {
		t.Errorf("Sample solution Score calculated as %d but should be %d", solution, 12)
	}
}
