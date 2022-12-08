package day7

import "testing"

func getSampleText() []string {
	sampleText := []string{
		"$ cd /",
		"$ ls",
		"dir a",
		"14848514 b.txt",
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

func TestIsFile(t *testing.T) {
	sampleText := getSampleText()

	if IsFile(sampleText[0]) == true {
		t.Error()
	}

	if IsFile(sampleText[1]) == true {
		t.Error()
	}

	if IsFile(sampleText[2]) == true {
		t.Error()
	}

	if IsFile(sampleText[3]) == false {
		t.Error()
	}
}

func TestIsDir(t *testing.T) {
	sampleText := getSampleText()

	if IsDir(sampleText[0]) == true {
		t.Error()
	}

	if IsDir(sampleText[1]) == true {
		t.Error()
	}

	if IsDir(sampleText[2]) == false {
		t.Error()
	}

	if IsDir(sampleText[3]) == true {
		t.Error()
	}
}
