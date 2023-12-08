package unit

import "testing"

func TestAdd(t *testing.T) {

	var a, b = 1, 2

	var execpet = 3

	if Add(a, b) != execpet {
		t.Errorf("Add(%d,%d) = %d,want %d", a, b, Add(a, b), execpet)
	}
}
