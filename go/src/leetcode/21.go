package leetcode

//ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

//MergeTwoLists ...
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	listM := new(ListNode)
	listScan := listM
	list1 := l1
	list2 := l2

	for {
		if list1 == nil {
			listScan.Next = list2
			break
		} else if list2 == nil {
			listScan.Next = list1
			break
		}

		if list1.Val <= list2.Val {
			listScan.Next = list1
			list1 = list1.Next
		} else {
			listScan.Next = list2
			list2 = list2.Next
		}

		listScan = listScan.Next
	}

	return listM.Next
}
