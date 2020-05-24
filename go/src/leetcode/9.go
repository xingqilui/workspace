package leetcode

//IsPalindrome 偷懒解法，直接使用问题7的方法，将整数反转，然后与原值比较是否相等
func IsPalindrome(x int) bool {
	var ret int64

	for z := int64(x); ; z /= 10 {
		ret += z % 10

		if z == z%10 {
			break
		}

		ret *= 10
	}

	if (ret > 2147483647) || (ret < -2147483648) || (ret < 0) {
		return false
	}

	if int(ret) == x {
		return true
	}

	return false
}
