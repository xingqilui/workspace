package leetcode

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	for k, v := range []struct {
		nums   []int
		target int
		exp    []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{0, 1, 2, 3, 4}, 7, []int{3, 4}},
		{[]int{-1, -2, 0, 1, 2}, 0, []int{0, 3}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	} {
		if cap := TwoSum(v.nums, v.target); true != reflect.DeepEqual(cap, v.exp) {
			t.Errorf("Testcase %d: expected %d but instead got %d!", k, v.exp, cap)
		}

		if cap := TwoSum1(v.nums, v.target); true != reflect.DeepEqual(cap, v.exp) {
			t.Errorf("Testcase %d: expected %d but instead got %d!", k, v.exp, cap)
		}

	}
}

func Test7(t *testing.T) {
	for k, v := range []struct {
		num int
		exp int
	}{
		{123, 321},
		{-123, -321},
		{1, 1},
		{120, 21},
		{-120, -21},
		{-120, -21},
	} {
		if cap := Reverse(v.num); cap != v.exp {
			t.Errorf("Testcase %d: expected %d but instead got %d!", k, v.exp, cap)
		}
	}
}

func Test9(t *testing.T) {
	for k, v := range []struct {
		num int
		exp bool
	}{
		{121, true},
		{-121, false},
		{10, false},
		{0, true},
		{2147483647, false},
		{-2147483648, false},
	} {
		if cap := IsPalindrome(v.num); cap != v.exp {
			t.Errorf("Testcase %d: expected %t but instead got %t!", k, v.exp, cap)
		}
	}
}

func Test13(t *testing.T) {
	for k, v := range []struct {
		rom string
		exp int
	}{
		{"I", 1},
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
		{"MMMCMXCIX", 3999},
	} {
		if cap := RomanToInt(v.rom); cap != v.exp {
			t.Errorf("Testcase %d: expected %d but instead got %d!", k, v.exp, cap)
		}
	}
}

func Test14(t *testing.T) {
	for k, v := range []struct {
		input []string
		exp   string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"abcd", "abab", "abab"}, "ab"},
		{[]string{"ab", "ab", "ab"}, "ab"},
		{[]string{"", "ab", "ab"}, ""},
		{[]string{"ab", "", ""}, ""},
		{[]string{}, ""},
	} {
		if cap := LongestCommonPrefix(v.input); cap != v.exp {
			t.Errorf("Testcase %d: expected %s but instead got %s!", k, v.exp, cap)
		}
	}
}

func Test20(t *testing.T) {
	for k, v := range []struct {
		input string
		exp   bool
	}{
		{"", true},
		{"()", true},
		{"[](){}", true},
		{"[({})]", true},
		{"{()[]}", true},
		{"{{{}}}", true},
		{"{{{[]()}}}", true},
		{"(]", false},
		{"([)]", false},
		{"([)", false},
		{"()]", false},
	} {
		if cap := IsValid(v.input); cap != v.exp {
			t.Errorf("Testcase %d: expected %t but instead got %t!", k, v.exp, cap)
		}

		if cap := IsValid2(v.input); cap != v.exp {
			t.Errorf("Testcase %d: expected %t but instead got %t!", k, v.exp, cap)
		}
	}

	for k, v := range []struct {
		input string
		exp   bool
	}{
		{"aaa", true},
		{"(abc)", true},
		{"abc[abc]abc(abc){abc}abc", true},
		{"[(abc{abc})]", true},
		{"{()abc[abc]}", true},
		{"{abc{{abc}}}", true},
		{"{{abc{[]abc()}abc}}", true},
		{"(]", false},
		{"([)]", false},
		{"([)", false},
		{"()]", false},
	} {
		if cap := IsValid3(v.input); cap != v.exp {
			t.Errorf("Testcase %d: expected %t but instead got %t!", k, v.exp, cap)
		}
		if cap := IsValid4(v.input); cap != v.exp {
			t.Errorf("Testcase %d: expected %t but instead got %t!", k, v.exp, cap)
		}

	}

}

func Test26(t *testing.T) {
	for k, v := range []struct {
		input []int
		exp   int
	}{
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
		{[]int{}, 0},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1}, 1},
		{[]int{1, 1, 2}, 2},
		{[]int{5, 5, 5, 4, 3, 3, 2, 1}, 5},
		{[]int{-1, 0, 1}, 3},
	} {
		if cap := RemoveDuplicates(v.input); cap != v.exp {
			t.Errorf("Testcase %d: expected %d but instead got %d!", k, v.exp, cap)
		}
	}
}

func Test27(t *testing.T) {
	for k, v := range []struct {
		input  []int
		input2 int
		exp    int
	}{
		{[]int{3, 2, 2, 3}, 3, 2},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
		{[]int{}, 1, 0},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1}, 1, 0},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1}, 2, 8},
		{[]int{1, 1, 1, 2, 3, 3, 4}, 1, 4},
		{[]int{-1, 0, 1}, -1, 2},
	} {
		if cap := RemoveElement(v.input, v.input2); cap != v.exp {
			t.Errorf("Testcase %d: expected %d but instead got %d!", k, v.exp, cap)
		}
	}
}

func Test28(t *testing.T) {
	for k, v := range []struct {
		input  string
		input2 string
		exp    int
	}{
		{"hello", "ll", 2},
		{"aaaaa", "bba", -1},
		{"aaaaa", "", 0},
		{"aaaaa", "a", 0},
		{"aaaaa", "aaa", 0},
		{"", "bba", -1},
		{"", "", 0},
		{"abcde", "c", 2},
		{"abcde", "abc", 0},
		{"abcde", "cde", 2},
		{"abcde", "def", -1},
		{"a", "a", 0},
		{"mississippi", "a", -1},
	} {
		if cap := StrStr(v.input, v.input2); cap != v.exp {
			t.Errorf("Testcase %d: expected %d but instead got %d!", k, v.exp, cap)
		}
	}
}

func Test35(t *testing.T) {
	for k, v := range []struct {
		input  []int
		input2 int
		exp    int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 5, 6}, 0, 0},
		{[]int{1, 1, 1, 1}, 1, 0},
		{[]int{1, 1, 1, 1}, 2, 4},
		{[]int{}, 2, 0},
		{[]int{1}, 0, 0},
		{[]int{1}, 1, 0},
		{[]int{1}, 2, 1},

		{[]int{1}, 2, 1},
	} {
		if cap := searchInsert(v.input, v.input2); cap != v.exp {
			t.Errorf("Testcase %d: expected %d but instead got %d!", k, v.exp, cap)
		}
	}
}

func Test53(t *testing.T) {
	for k, v := range []struct {
		input []int
		exp   int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{-1, 0, 1}, 1},
		{[]int{1, 2, 3}, 6},
		{[]int{10, -1, 0, 1, 10}, 20},
		{[]int{-1}, -1},
	} {
		if cap := maxSubArray(v.input); cap != v.exp {
			t.Errorf("Testcase %d: expected %d but instead got %d!", k, v.exp, cap)
		}
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
