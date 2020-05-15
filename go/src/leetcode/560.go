package leetcode

//SubarraySum560 ...
func SubarraySum560(nums []int, k int) int {
	var ret int

	for i := range nums {
		sum := 0

		for end := i; end < len(nums); end++ {
			sum += nums[end]

			if sum == k {
				ret++
			} else {
				continue
			}
		}
	}

	return ret
}

//SubarraySum560_2 前缀和算法
func SubarraySum560_2(nums []int, k int) int {
	var ret, sum int
	var m = map[int]int{0: 1}

	for _, v := range nums {
		sum += v

		if _, ok := m[sum-k]; ok == true {
			ret += m[sum-k]
		}

		m[sum]++
	}

	return ret
}

//[3,4,7,2,-3]
