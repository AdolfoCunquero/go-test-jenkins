package test

import (
	"testing"
)

func Sum(x int, y int) int {
	return x + y
}

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}

	total2 := Sum(15, 5)
	if total2 != 20 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total2, 10)
	}
}

func TestSum2(t *testing.T) {
	total := Sum(6, 5)
	if total != 11 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}

	total2 := Sum(5, 15)
	if total2 != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total2, 10)
	}
}
