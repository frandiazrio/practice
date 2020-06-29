package main

import "testing"

func TestSum2(t *testing.T) {
	total := Sum(0, 0)
	if total != 0 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}

	total = Sum(-1, 5)
	if total != 4 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}
