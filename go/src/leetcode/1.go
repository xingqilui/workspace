package leetcode

//TwoSum 使用map实现，一次遍历填充map，一次遍历检索答案
func TwoSum(nums []int, target int) []int {
	var nummap = map[int]int{}

	for k, v := range nums {
		nummap[v] = k
	}

	for k, v := range nums {
		if val, ok := nummap[target-v]; true == ok && k != val {
			return []int{k, val}
		}
	}

	return []int{0, 0}
}

//TwoSum1 使用map一次遍历
func TwoSum1(nums []int, target int) []int {
	var nummap = map[int]int{}

	for k, v := range nums {
		if val, ok := nummap[target-v]; true == ok && k != val {
			return []int{val, k}
		}
		nummap[v] = k
	}

	return []int{0, 0}
}
