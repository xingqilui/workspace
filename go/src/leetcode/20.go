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

//IsValid2 使用切片模拟栈操作
func IsValid2(s string) bool {
	var m = map[byte]byte{'(': ')', '[': ']', '{': '}'}
	var stack = []byte{}
	for i := range s {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else {
			top := len(stack)

			if top <= 0 {
				return false
			}

			if m[stack[top-1]] == s[i] {
				stack = stack[:top-1]
			} else {
				return false
			}
		}
	}

	return (0 == len(stack))
}

//IsValid3 在括号之间加一些其他字符时，也能正常识别并匹配开闭区间
func IsValid3(s string) bool {
	var left = map[byte]byte{'(': ')', '[': ']', '{': '}'}
	var right = map[byte]byte{')': '(', ']': '[', '}': '{'}
	var stack = make([]byte, 0)

	for i := range s {
		if _, ok := left[s[i]]; true == ok {
			stack = append(stack, s[i])
		} else if _, ok := right[s[i]]; true == ok {
			top := len(stack)

			if top <= 0 {
				return false
			}

			if left[stack[top-1]] == s[i] {
				stack = stack[:top-1]
			} else {
				return false
			}
		} else {
			continue
		}
	}

	return (0 == len(stack))
}

//IsValid4 使用匿名函数实现栈的操作集合
func IsValid4(s string) bool {
	var left = map[byte]byte{'(': ')', '[': ']', '{': '}'}
	var right = map[byte]byte{')': '(', ']': '[', '}': '{'}
	var stack = make([]byte, 0)

	//栈的常规操作：入栈、出栈、取栈顶
	empty := func(stack *[]byte) bool {
		return 0 == len(*stack)
	}

	top := func(stack *[]byte) byte {
		if empty(stack) {
			return 0
		}

		return (*stack)[len(*stack)-1]
	}

	push := func(stack *[]byte, item byte) {
		*stack = append(*stack, item)
	}

	pop := func(stack *[]byte) {
		if empty(stack) {
			return
		}

		*stack = (*stack)[:len(*stack)-1]
	}

	for i := range s {
		if _, ok := left[s[i]]; true == ok {
			push(&stack, s[i])
		} else if _, ok := right[s[i]]; true == ok {
			if empty(&stack) {
				return false
			}

			if left[top(&stack)] == s[i] {
				pop(&stack)
			} else {
				return false
			}
		} else {
			continue
		}
	}

	return (0 == len(stack))
}
