package file_input

import (
	"testing"
)

const test_file = "./file_input_sample.txt"

func TestFirstStringLine(t *testing.T) {
	file_lines := Read_file(test_file)
	if file_lines[0] != "1000" {
		t.Error()
	}
}

func TestEmptyStringLine(t *testing.T) {
	file_lines := Read_file(test_file)
	content := file_lines[3]
	if content != "" {
		t.Errorf("Found %s instead of empty string", content)
	}
}

func TestTrailingEmptyLineIgnored(t *testing.T) {
	file_lines := Read_file(test_file)
	content := file_lines[len(file_lines)-1]
	if content != "10000" {
		t.Errorf("Found %s instead of expected string", content)
	}
}
