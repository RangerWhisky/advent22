package utils

import "testing"

func TestCreateBoolMap(t *testing.T) {
	boolmap := InitialiseBoolMap(10, 10)

	if boolmap.countTrue != 0 {
		t.Error()
	}
}
