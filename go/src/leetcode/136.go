package leetcode

//SingleNumber ...
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

//SingleNumber1 官方解法：使用异或运算。a^a=0, a^0=a，异或满足结合律，所以两两异或全部消为0，最终出现一次的数与0异或等于本身
func SingleNumber1(nums []int) int {
	var ret int
	for _, v := range nums {
		ret ^= v
	}

	return ret
}
