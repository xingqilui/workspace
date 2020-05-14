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
	var e1 = map[string]int{}
	var e2 = map[string]int{"身高": 120, "体重": 20}
	var e3 = make(map[string]int)

	e4 := map[string]int{}
	e5 := map[string]int{"身高": 140, "体重": 40}
	e6 := make(map[string]int)

	fmt.Println("e:", e1, e2, e3, e4, e5, e6, e2["身高"], e5["体重"])

	val, ok := e2["身高"]
	fmt.Println(val, ok)

	delete(e2, "身高")
	val, ok = e2["身高"]
	fmt.Println(val, ok)

	//结构体初始化
	type F struct {
		name string
		sex  bool
		age  int
	}

	var f1 F
	var f2 = F{"bbb", true, 20}
	var f3 = F{name: "ccc", sex: true, age: 30}
	f4 := F{}
	f5 := F{"ddd", true, 20}
	f6 := F{name: "eee", sex: true, age: 30}

	f1.name = "aaa"
	f1.sex = true
	f1.age = 30

	fmt.Println("f:", f1, f2, f3, f4, f5, f6)

	//结构体数组
	var g1 [3]F
	g1[0].name = "a"
	g1[1].name = "b"
	g1[2].name = "c"

	var g2 = [3]F{
		{"aa", true, 10},
		{"bb", true, 20},
		{"cc", true, 30},
	}

	var g3 = []F{
		{"aaa", true, 40},
		{"bbb", true, 50},
		{"ccc", true, 60},
	}

	g4 := []F{
		{"aaaa", true, 40},
		{"bbbb", true, 50},
		{"cccc", true, 60},
	}

	fmt.Println("g:", g1, g2, g3, g4)

	//结构体集合
	var h1 = map[string]F{}
	var h2 = make(map[string]F)
	h3 := make(map[string]F)
	h4 := map[string]F{}

	h1["a"] = F{"a", true, 10}
	h2["b"] = F{"b", true, 20}
	h3["c"] = F{"c", true, 30}
	h4["d"] = F{"d", true, 40}

	fmt.Println("h:", h1, h2, h3, h4)

	//结构体多成员类型初始化
	type I struct {
		id   [3]int
		info []string
		city []map[string]int
	}

	var i1 = I{}
	var i2 = I{[3]int{1, 2, 3}, []string{"a", "b", "c"}, []map[string]int{{"1": 1}, {"2": 2}}}
	var i3 = I{id: [3]int{1, 2, 3}, info: []string{"d", "e", "f"}, city: []map[string]int{{"3": 3}, {"4": 4}}}

	fmt.Println("i:", i1, i2, i3)

}
