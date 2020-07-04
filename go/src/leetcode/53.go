package leetcode

func maxSubArray(nums []int) int {
	var max int

	if len(nums) != 0 {
		max = nums[0]
	}

	for i := range nums {
		sum := 0

		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum > max {
				max = sum
			}
		}
	}

	return max
}

/*

# 笔记
[toc]

## 解题思路1
暴力解法，使用两层循环，遍历所有成员加和的情况，复杂度为$n^2$。

```
func maxSubArray(nums []int) int {
	var max int

	if len(nums) != 0 {
		max = nums[0]
	}

	for i := range nums {
		sum := 0

		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum > max {
				max = sum
			}
		}
	}

	return max
}
```

## 解题思路2

```

```

## 遇到的问题
1. 当输入为[-1]时，返回了0。
	> max的初始值写成0了，因为输入可以为正也可以为负，所以当nums成员非空时，max初始化为第一个成员。

## 知识点


*/
