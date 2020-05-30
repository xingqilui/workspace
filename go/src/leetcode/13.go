package leetcode

//RomanToInt ...
func RomanToInt(s string) int {
	var pre, num int
	var roman = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for k := range s {
		if roman[s[k]] > pre {
			num += roman[s[k]] - pre - pre
		} else {
			num += roman[s[k]]
		}

		pre = roman[s[k]]
	}

	return num
}
