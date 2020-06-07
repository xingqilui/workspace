package leetcode

//IsValid 使用数组模拟栈的操作，进行入栈和出栈，匹配括号的开闭区间
func IsValid(s string) bool {
	var m = map[byte]byte{'(': ')', '[': ']', '{': '}'}
	var stack = make([]byte, len(s))
	var index = -1

	for i := range s {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			index++
			stack[index] = s[i]
		} else {
			if index < 0 {
				return false
			}

			if m[stack[index]] == s[i] {
				stack[index] = 0
				index--
			} else {
				return false
			}
		}
	}

	return (index == -1)
}
