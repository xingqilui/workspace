package leetcode

//RemoveElement ...
func RemoveElement(nums []int, val int) int {
	var i, j int

	if len(nums) == 0 {
		return 0
	}

	for i = range nums {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
	}

	return j
}

/*

# 笔记

## 解题思路1
与26题解题思路一样，使用两个游标。

```go
func RemoveElement(nums []int, val int) int {
	var i, j int

	if len(nums) == 0 {
		return 0
	}

	for i = range nums {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
	}

	return j
}
```

## 解题思路2

```go

```

## 知识点
1. 数组作为函数入参，是传值还是传指针？
	> * 数组作为函数参数是值传递，函数内改变不会改变实参的值
	> * 切片作为函数参数是地址传递，函数内改变会改变实参的值

*/
