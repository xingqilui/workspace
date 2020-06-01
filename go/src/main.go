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

}
