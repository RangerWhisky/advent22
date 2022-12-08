package day7

import "testing"

func TestNavigateFilesystem(t *testing.T) {
	fs := InitialiseFilesystem()

	currentDir := PrintWorkingDirectory(&fs)

	if currentDir != "/" {
		t.Errorf("CurrentDir not properly initialised, reporting %q", currentDir)
	}

	ChangeDir(&fs, "a")

	currentDir = PrintWorkingDirectory(&fs)

	if currentDir != "a" {
		t.Errorf("CurrentDir not changed, reporting %q", currentDir)
	}

	ChangeDir(&fs, "..")

	currentDir = PrintWorkingDirectory(&fs)

	if currentDir != "/" {
		t.Errorf("CurrentDir not changed, reporting %q", currentDir)
	}

}

func TestSaveDirectory(t *testing.T) {

	sampleText := [][]byte{
		[]byte("dir a"),
		[]byte("14848514 b.txt"),
	}

	dir := ParseDirectory(sampleText)

	fs := InitialiseFilesystem()

	SaveDirectory(&fs, dir)

	savedDir := GetDirectory(&fs, "/")
	if GetSize(&savedDir) != 14848514 {
		t.Errorf("savedDir %q", savedDir)
	}
}
