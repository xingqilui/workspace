package leetcode

//Reverse 偷懒解法，用int64接一下入参，返回时判断是否越int边界
func Reverse(x int) int {
	var ret int64

	for z := int64(x); ; z /= 10 {
		ret += z % 10

		if z == z%10 {
			break
		}

		ret *= 10
	}

	if (ret > 2147483647) || (ret < -2147483648) {
		return 0
	}

	return int(ret)
}
