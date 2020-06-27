package leetcode

//StrStr ...
func StrStr(haystack string, needle string) int {
	index := -1
	needleSize := len(needle)
	haystackSize := len(haystack)

	if haystackSize == 0 && needleSize == 0 {
		return 0
	}

	if haystackSize == 0 {
		return -1
	}

	isSliceEqual := func(a string, b string) bool {
		if len(a) != len(b) {
			return false
		}

		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}

		return true
	}

	for i := range haystack {
		if i+needleSize > haystackSize {
			return -1
		}

		if true == isSliceEqual(haystack[i:i+needleSize], needle) {
			index = i
			break
		}
	}

	return index
}

/*

# 笔记

## 解题思路1
遍历源串，使用字串和源串的【当前索引+字串长度】的切片和子串判断，如果相等，立即返回当前索引。

```go
func StrStr(haystack string, needle string) int {
	index := -1
	needleSize := len(needle)
	haystackSize := len(haystack)

	if haystackSize == 0 && needleSize == 0 {
		return 0
	}

	if haystackSize == 0 {
		return -1
	}

	isSliceEqual := func(a string, b string) bool {
		if len(a) != len(b) {
			return false
		}

		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}

		return true
	}

	for i := range haystack {
		if i+needleSize > haystackSize {
			return -1
		}

		if true == isSliceEqual(haystack[i:i+needleSize], needle) {
			index = i
			break
		}
	}

	return index
}
```

## 解题思路2

```go

```

## 遇到的问题
1. index初始化时应该为-1，在needle完全匹配不到haystack时，要返回默认index值。
2. 遍历中判断边界值，使用>号，而不是>=号，在haystack和needle完全相同时，要返回0.
3. 当源串为空时，子串也为空则返回0，子串非空则返回-1。
4. 当源串非空时，子串为空时则返回0。

## 知识点
1. go中对数组的判断，可以直接使用==号，数组的长度和成员完全一致，则返回true，否则返回false。
2. go中对切片的判断，不能直接使用==号，需要自己循环遍历，或者使用DeepEqual方法。
	> 官方的reflect包中有个DeepEqual方法，可以用来判断任意x和y是否相等。相同类型的两个值可能相等，不同类型的值永远不会相等。
3. 搜索文本中的字串，使用KMP算法。
*/
