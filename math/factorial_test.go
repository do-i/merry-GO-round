package math

import "testing"

func FactorialOf_valid(t *testing.T) {
	result := FactorialOf(5)
	if result != 200 {
		t.Error("expected 200, got ", result)
	}
}