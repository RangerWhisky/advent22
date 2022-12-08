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
