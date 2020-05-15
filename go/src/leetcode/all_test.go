package leetcode

import (
	"testing"
)

func TestSum(t *testing.T) {
	num1, num2, exp := 1, 2, 3

	act := Sum(num1, num2)
	if act != exp {
		t.Errorf("Expected the sum of %d + %d to be %d but instead got %d!", num1, num2, exp, act)
	}
}

func Test136(t *testing.T) {
	for k, v := range []struct {
		nums []int
		exp  int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
		{[]int{2, 1, 2, 1, 0}, 0},
		{[]int{1, 1, 2, 2, 0}, 0},
		{[]int{1, 1, 987654321, 2, 2}, 987654321},
	} {
		if cap := SingleNumber(v.nums); cap != v.exp {
			t.Errorf("Testcase %d: expected %d but instead got %d!", k, v.exp, cap)
		}

		if cap := SingleNumber1(v.nums); cap != v.exp {
			t.Errorf("Testcase %d: expected %d but instead got %d!", k, v.exp, cap)
		}
	}
}

func Test137(t *testing.T) {
	for k, v := range []struct {
		nums []int
		exp  int
	}{
		{[]int{2, 2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2, 1, 2}, 4},
		{[]int{2, 1, 2, 1, 2, 1, 0}, 0},

		{[]int{1, 1, 1, 2, 2, 2, 0}, 0},
		{[]int{1, 1, 1, 987654321, 2, 2, 2}, 987654321},
	} {
		if cap := SingleNumber137(v.nums); cap != v.exp {
			t.Errorf("Testcase %d: expected %d but instead got %d!", k, v.exp, cap)
		}

		if cap := SingleNumber137_1(v.nums); cap != v.exp {
			t.Errorf("Testcase %d: expected %d but instead got %d!", k, v.exp, cap)
		}
	}
}

func Test560(t *testing.T) {
	for k, v := range []struct {
		nums []int
		k    int
		exp  int
	}{
		{[]int{1, 1, 1}, 2, 2},
		{[]int{1, 2, 3}, 5, 1},
		{[]int{1, 2, 3}, 4, 0},
		{[]int{1, 1, 1, 1, 1}, 2, 4},
		{[]int{1}, 1, 1},
		{[]int{-1, 1, 2, 3, -2, -3}, 0, 3},
		{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 0, 55},
	} {
		if cap := SubarraySum560(v.nums, v.k); cap != v.exp {
			t.Errorf("Testcase %d: expected %d but instead got %d!", k, v.exp, cap)
		}

		if cap := SubarraySum560_2(v.nums, v.k); cap != v.exp {
			t.Errorf("Testcase %d: expected %d but instead got %d!", k, v.exp, cap)
		}
	}
}
