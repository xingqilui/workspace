package leetcode

//RemoveDuplicates ...
func RemoveDuplicates(nums []int) int {
	var i, j int

	if len(nums) == 0 {
		return 0
	}

	for i = range nums {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}

	return j + 1
}

/*

# 笔记

## 解题思路1
设置两个游标，一个是i，遍历整个数组。一个是j，用来更新数组值。
使用一次循环，i和j起点都是索引0，比较如果两个索引相等，则i自增，否则j+1，然后将索引i的值填入索引j。

```go
func RemoveDuplicates(nums []int) int {
	var i, j int

	if len(nums) == 0 {
		return 0
	}

	for i = range nums {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}

	return j + 1
}
```

## 解题思路2


## 遇到的问题



*/
