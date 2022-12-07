package day7

import "testing"

func TestGetSimplestDirectorySize(t *testing.T) {

	sampleText := [][]byte{[]byte("14848514 b.txt")}

	size := GetDirectorySize(sampleText)

	if size != 14848514 {
		t.Error()
	}
}
