package main

import "fmt"

func main() {
	//变量初始化
	var a1 int
	var a2 = 10
	a3 := 30

	var a4, a5 int
	var a6, a7 = 60, 70
	a8, a9 := 80, 90

	fmt.Println(a1, a2, a3, a4, a5, a6, a7, a8, a9)

	//常量初始化
	const b1 int = 100
	const b2 = 200

	//多常量初始化
	const b3, b4, b5, b6 = 300, "400", false, 600.66

	//枚举
	const (
		b7  = 700.77
		b8  = "800"
		b9  = 900
		b10 = iota
		b11 = b10
		b12 = iota
	)

	fmt.Println(b1, b2, b3, b4, b5, b6, b7, b8, b9, b10, b11, b12)

	//数组初始化
	var c1 [3]int
	var c2 = [3]int{21, 22, 23}
	var c3 = [...]int{31, 32, 33}

	c4 := [3]int{}
	c5 := [3]int{51, 52, 53}
	c6 := [...]int{61, 62, 63}

	fmt.Println(c1, c2, c3, c4, c5, c6)

	//切片初始化
	var d1 []int
	var d2 = []int{21, 22, 23}
	var d3 = make([]int, 3)

	d4 := []int{}
	d5 := []int{51, 52, 53}
	d6 := make([]int, 3)

	d7 := d2[:]
	d8 := d2[0:]
	d9 := d2[:2]
	d10 := d2[0:2]

	fmt.Println(d1, d2, d3, d4, d5, d6, d7, d8, d9, d10)

	d1 = append(d2, 1, 2, 3, 4, 5)

	fmt.Println(d1, len(d1), cap(d1))

	//集合初始化
	var e1 map[string]int
	var e2 = map[string]int{"身高": 120, "体重": 20}

	e3 := map[string]int{}
	e4 := map[string]int{"身高": 140, "体重": 40}

	fmt.Println(e1, e2, e3, e4, e2["身高"], e4["体重"])

	val, ok := e2["身高"]
	fmt.Println(val, ok)

	delete(e2, "身高")
	val, ok = e2["身高"]
	fmt.Println(val, ok)
}
