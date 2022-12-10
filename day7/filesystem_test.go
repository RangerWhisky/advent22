package day7

import "testing"

func TestNavigateFilesystem(t *testing.T) {
	fs := InitialiseEmptyFilesystem()

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

	fs := InitialiseEmptyFilesystem()

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

	fs := InitialiseEmptyFilesystem()

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

func TestGetDirContentIndices(t *testing.T) {
	sample := []string{"$ cd /",
		"$ ls",
		"dir a",
		"14848514 b.txt",
		"8504156 c.dat",
		"$ cd a",
		"$ ls",
		"29116 f",
		"2557 g",
		"62596 h.lst"}

	start, end := GetDirContentIndices(sample, 1)

	if start != 2 {
		t.Errorf("Start is wrong - %d", start)
	}
	if end != 5 {
		t.Errorf("End is wrong - %d", end)
	}
}

func TestInitialiseFileSystem(t *testing.T) {
	sample := []string{"$ cd /",
		"$ ls",
		"dir a",
		"14848514 b.txt",
		"8504156 c.dat",
		"$ cd a",
		"$ ls",
		"29116 f",
		"2557 g",
		"62596 h.lst"}

	fs := InitialiseFilesystem(sample)
	size := GetSizeOnDisk(&fs, "a")

	if size != (29116 + 2557 + 62596) {
		t.Errorf("Size is %d for filesystem %q", size, fs)
	}
}
