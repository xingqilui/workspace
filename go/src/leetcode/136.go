package leetcode

//SingleNumber done
func SingleNumber(nums []int) int {
	var m = map[int]int{}

	for _, v := range nums {
		m[v]++
	}

	for k, v := range m {
		if 1 == v {
			return k
		}
	}

	return 0
}
