package tools

import (
	"testing"
)

type Input struct {
	num1 int
	num2 int
	exp  int
}

func TestSum(t *testing.T) {

	input := []Input{
		{1, 2, 3},
		{10, 20, 30},
	}

	for _, v := range input {
		act := Sum(v.num1, v.num2)
		if act != v.exp {
			t.Errorf("Expected the sum of %d + %d to be %d but instead got %d!", v.num1, v.num2, v.exp, act)
		}
	}
}
