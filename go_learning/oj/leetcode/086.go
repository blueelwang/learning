package leetcode

import (
	"fmt"
)

func Partition() {
	n1, n2, n3, n4, n5, n6, n7 := &ListNode{2, nil}, &ListNode{8, nil}, &ListNode{3, nil}, &ListNode{6, nil}, &ListNode{7, nil}, &ListNode{4, nil}, &ListNode{1, nil}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6
	n6.Next = n7
	printListNode(n1)
	head := partition(n1, 3)
	printListNode(head)
}

func partition(head *ListNode, x int) *ListNode {
	list1, list2 := &ListNode{0, nil}, &ListNode{0, nil}
	tail1, tail2 := list1, list2

	for nil != head {
		if head.Val < x {
			tail1.Next = head
			tail1 = tail1.Next
			head = head.Next
			tail1.Next = nil
		} else {
			tail2.Next = head
			tail2 = tail2.Next
			head = head.Next
			tail2.Next = nil
		}
	}

	tail1.Next = list2.Next
	return list1.Next
}

func printListNode(head *ListNode) {
	for ;nil != head; head = head.Next {
		fmt.Printf("%d -> ", head.Val)
	}
	fmt.Println("")
}
