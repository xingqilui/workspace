package main

import (
	"fmt"
	"tools"
)

func main() {
	fmt.Println("Hello world!", tools.Sum(100, 200))

	a := [...]int{1, 2, 3}
	b := a

	fmt.Println(a == b)

}
