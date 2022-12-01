package day1

import (
	"testing"
)

func TestHello(t *testing.T) {
	meaning := Meaning()
	if meaning != 42 {
		t.Errorf("Meaning is = %d, need 42", meaning)

	}
}
