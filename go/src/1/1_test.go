package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	num1, num2, exp := 1, 2, 4

	act := sum(num1, num2)
	if act != exp {
		t.Errorf("Expected the sum of %d + %d to be %d but instead got %d!", num1, num2, exp, act)
	}
}
