package math

import "testing"

func CombinationOf_valid(t *testing.T) {
	result := CombinationOf(5, 3)
	if result != 10 {
		t.Error("expected 10, got ", result)
	}
}