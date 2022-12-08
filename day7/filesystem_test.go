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

	sampleText := []string{
		"dir a",
		"14848514 b.txt",
	}

	dir := ParseDirectory(sampleText)

	fs := InitialiseFilesystem()

	SaveDirectory(&fs, dir)

	savedDir := GetDirectory(&fs, "/")
	if GetSize(&savedDir) != 14848514 {
		t.Errorf("savedDir %q", savedDir)
	}
}

func TestRecursiveDirectorySize(t *testing.T) {

	sampleRoot := []string{
		"dir a",
		"14848514 b.txt",
	}

	dir := ParseDirectory(sampleRoot)

	fs := InitialiseFilesystem()

	SaveDirectory(&fs, dir)

	ChangeDir(&fs, "a")

	sampleContents := []string{
		"11 c.txt",
	}

	a := ParseDirectory(sampleContents)
	SaveDirectory(&fs, a)

	size := GetSizeOnDisk(&fs, "/")
	if size != 14848525 {
		t.Errorf("size, %d, should be %d", size, 14848525)
	}
}
