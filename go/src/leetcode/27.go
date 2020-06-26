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

## 遇到的问题

## 思考的问题



*/
