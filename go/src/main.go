package main

import (
	"fmt"
	"tools"
)

func main() {
	fmt.Println("Hello world!", tools.Sum(100, 200))

	fmt.Println(123 % 10000)

	s := "hello"

	for k, v := range s {
		fmt.Printf("%d:%c\n", k, v)
	}
}
