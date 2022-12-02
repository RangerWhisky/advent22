package day2

import "testing"

func TestGetResponseScore(t *testing.T) {
	testInput := "A X"
	score := GetResponseScore(testInput)
	if score != 1 {
		t.Errorf("Response score is = %d, need 1", score)
	}
}
