package day7

import "testing"

func getSampleText() [][]byte {
	sampleText := [][]byte{
		[]byte("$ cd /"),
		[]byte("$ cd /"),
		[]byte("dir a"),
		[]byte("14848514 b.txt"),
	}

	return sampleText

}

func TestIsCmd(t *testing.T) {
	sampleText := getSampleText()

	if IsCmd(sampleText[0]) == false {
		t.Error()
	}

	if IsCmd(sampleText[1]) == false {
		t.Error()
	}

	if IsCmd(sampleText[2]) == true {
		t.Error()
	}

	if IsCmd(sampleText[3]) == true {
		t.Error()
	}
}
