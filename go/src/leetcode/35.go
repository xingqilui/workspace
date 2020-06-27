package leetcode

func searchInsert(nums []int, target int) int {
	var i int

	if 0 == len(nums) {
		return 0
	}

	for i = range nums {
		if nums[i] >= target {
			return i
		}
	}

	return i + 1
}

/*

# 笔记

## 解题思路1
遍历找到 >= 目标的值，就可以返回当前索引了。如果遍历到结尾仍然没有找打， 则说明目标是最大的，返回nums的len + 1。

```
func searchInsert(nums []int, target int) int {
	var i int

	if 0 == len(nums) {
		return 0
	}

	for i = range nums {
		if nums[i] >= target {
			return i
		}
	}

	return i + 1
}
```

## 解题思路2

```

```

## 遇到的问题
1. 注意目标值如果是最大的值，则返回源数组的长度+1。

## 知识点


*/
