package leetcode

// LongestCommonPrefix 暴力解法
func LongestCommonPrefix(strs []string) string {
	var minlen = 0xffffffff
	var minlenindex int

	if 0 == len(strs) {
		return ""
	}

	//找到最短的字符串
	for i := range strs {
		strlen := len(strs[i])

		if strlen < minlen {
			minlen = strlen
			minlenindex = i
		}
	}

	if 0 == minlen {
		return ""
	}

	for i := 0; i < minlen; i++ {
		for k := range strs {
			if strs[minlenindex][i] != strs[k][i] {
				return strs[minlenindex][:i]
			}
		}
	}

	return strs[minlenindex]
}
