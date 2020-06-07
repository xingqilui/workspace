package leetcode

//ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

//MergeTwoLists ...
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	// listAdd := func(list *ListNode, item *ListNode) {
	// 	item.Next = list.Next
	// 	list.Next = item
	// }

	// listDelHead := func(list *ListNode) {
	// 	list.Next = list.Next.Next
	// }

	// listReverse := func(list *ListNode) {
	// 	var listRet = new(ListNode)

	// 	for list.Next != nil {
	// 		item := list.Next
	// 		listDelHead(list)
	// 		listAdd(listRet, item)
	// 	}

	// 	list.Next = listRet.Next
	// }

	// listShow := func(list *ListNode) {
	// 	fmt.Printf("List show start at %p\n", list)

	// 	for item := list.Next; item != nil; item = item.Next {
	// 		fmt.Printf("Addr=%p, val=%d\n", item, item.Val)
	// 	}

	// 	fmt.Println("List show end.")
	// }

	// listReverse(list)
	// listShow(list)

	list := new(ListNode)
	item := list
	item1 := l1
	item2 := l2

	for {
		if item1.Next == nil {
			list.Next = item2
			break
		} else {
			list.Next = item1
			break
		}

		if item1.Val <= item2.Val {
			item = item1
			item1 = item1.Next
		} else {
			item = item2
			item2 = item2.Next
		}

		l1.Next = item.Next
		item.Next = list.Next
		list.Next = item
	}

	return list.Next
}
