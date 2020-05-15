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
