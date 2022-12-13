package utils

import "testing"

func TestCreateBoolMap(t *testing.T) {
	boolmap := InitialiseBoolMap(10, 10)

	if boolmap.countTrue != 0 {
		t.Error()
	}
}

func TestSetTrue(t *testing.T) {
	boolmap := InitialiseBoolMap(10, 10)

	if boolmap.countTrue != 0 {
		t.Error()
	}

	MarkMap(&boolmap, 5, 5)

	if boolmap.countTrue != 1 {
		t.Error()
	}

}
