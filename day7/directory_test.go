package day7

import "testing"

func TestGetSimplestDirectorySize(t *testing.T) {

	sampleText := []string{"14848514 b.txt"}
	directory := ParseDirectory(sampleText)

	size := GetSize(&directory)

	if size != 14848514 {
		t.Error()
	}
}

func TestParseDirectory(t *testing.T) {

	sampleText := []string{
		"dir a",
		"14848514 b.txt",
	}

	dir := ParseDirectory(sampleText)

	if dir.size != 14848514 {
		t.Error()
	}

	if dir.subdirs[0] != "a" {
		t.Error()
	}
}
