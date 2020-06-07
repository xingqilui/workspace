package tools

import "fmt"

//Sum ...
func Sum(num1 int, num2 int) int {
	return num1 + num2
}

//SList ...
type SList struct {
	Next   *SListNode
	Length int
}

//SListNode ...
type SListNode struct {
	Next *SListNode
	Data int
}

//SListInit ...
func (list *SList) SListInit(nums []int) {

	for i := range nums {
		item := new(SListNode)
		item.Data = nums[i]
		list.SLinkAddHead(item)
	}
}

//SLinkAddHead ..
func (list *SList) SLinkAddHead(item *SListNode) {
	item.Next = list.Next
	list.Next = item
	list.Length++
}

//SLinkAddNode ..
func (list *SList) SLinkAddNode(src *SListNode, item *SListNode) {
	item.Next = src.Next
	src.Next = item
	list.Length++
}

//SLinkDelHead ...
func (list *SList) SLinkDelHead() {
	list.Next = list.Next.Next
	list.Length--
}

//SLinkReverse ...
func (list *SList) SLinkReverse() {
	var listRet SList

	for list.Next != nil {
		item := list.Next
		list.SLinkDelHead()
		listRet.SLinkAddHead(item)
	}

	list.Next = listRet.Next
}

//SLinkMerge ...
// func (list *SList) SLinkMerge(l1 *SList, l2 *SList) *SList {

// }

//SLinkShow ..
func (list *SList) SLinkShow() {
	fmt.Printf("List show start at %p\n", list)

	for item := list.Next; item != nil; item = item.Next {
		fmt.Printf("Addr=%p, val=%d\n", item, item.Data)
	}

	fmt.Println("List show end.")
}
