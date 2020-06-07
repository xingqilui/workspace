package main

import (
	"fmt"
	"tools"
)

func main() {
	fmt.Println("Hello world!", tools.Sum(100, 200))

	var str1 = ""
	var str2 = "abc"
	str1 = str2[:1]

	fmt.Println(str1)
	fmt.Println(len(str1))

	var a byte
	fmt.Println(a)

	var stack = []byte{'a', 'b'}

	//栈的常规操作：入栈、出栈、取栈顶
	empty := func(stack *[]byte) bool {
		return 0 == len(*stack)
	}

	// top := func(stack []byte) byte {
	// 	if empty(stack) {
	// 		return 0
	// 	}

	// 	return stack[len(stack)-1]
	// }

	push := func(stack *[]byte, item byte) {
		*stack = append(*stack, item)
	}

	// pop := func(stack []byte) {
	// 	if empty(stack) {
	// 		return
	// 	}

	// 	stack = stack[:len(stack)-1]
	// }

	push(&stack, 'a')
	fmt.Println(stack)
	fmt.Println(empty(&stack))
}
