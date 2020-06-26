package main

import (
	"fmt"
	"tools"
)

//Person ...
type Person struct {
	name string
	age  int
}

//Student ...
type Student struct {
	Person
	id int
}

//Hello ...
func (person *Person) Hello() {
	fmt.Printf("Hello, My name is %s, I`m %d years old.\n", person.name, person.age)
}

//Growup ...
func (person *Person) Growup() {
	person.age++
}

//Say ...
func (student *Student) Say() {
	fmt.Printf("My Id is %d.\n", student.id)
}

//ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Println("Hello world!", tools.Sum(100, 200))

	var chenlei = Person{"chenlei", 30}

	fmt.Println(chenlei)
	chenlei.Hello()
	chenlei.Growup()
	chenlei.Growup()
	chenlei.Growup()
	chenlei.Hello()

	var liuyueshu = Student{Person{"Liuyueshu", 30}, 10010}
	liuyueshu.Hello()
	liuyueshu.Say()

	var list = new(ListNode)

	var item1 = new(ListNode)
	item1.Val = 100
	item1.Next = nil

	var item2 = new(ListNode)
	item2.Val = 200
	item2.Next = nil

	listAdd := func(list *ListNode, item *ListNode) {
		item.Next = list.Next
		list.Next = item
	}

	listDelHead := func(list *ListNode) {
		list.Next = list.Next.Next
	}

	listReverse := func(list *ListNode) {
		var listRet = new(ListNode)

		for list.Next != nil {
			item := list.Next
			listDelHead(list)
			listAdd(listRet, item)
		}

		list.Next = listRet.Next
	}

	listShow := func(list *ListNode) {
		fmt.Printf("List show start at %p\n", list)

		for item := list; item != nil; item = item.Next {
			fmt.Printf("Addr=%p, val=%d\n", item, item.Val)
		}

		fmt.Println("List show end.")
	}

	listAdd(list, item1)
	listAdd(list, item2)
	listShow(list)

	listReverse(list)
	listShow(list)

	var list1 = new(tools.ListNode)
	var list2 = new(tools.ListNode)

	tools.ListInit(list1, []int{5, 4, 3, 2, 1})
	tools.ListShow(list1)
	tools.ListInit(list2, []int{10, 9, 8, 7, 6, 5, 4, 3})
	tools.ListShow(list2)

	listM := tools.MergeTwoLists(list1, list2)
	tools.ListShow(listM)

	// var slist = new(tools.SList)
	// slist.SListInit([]int{1, 2, 3, 4, 5})
	// slist.SLinkShow()
	// slist.SLinkReverse()
	// slist.SLinkShow()

	fmt.Println("-----26-----")

	arrayModify := func(nums []int) {
		nums[0] = 100
	}

	arrayNums := []int{1, 2, 3}
	arrayModify(arrayNums)
	fmt.Println(arrayNums[0])

}
