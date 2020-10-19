package leetcode


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	head := &ListNode{
		Val : 0,
		Next : nil,
	}
	tail := head

	for {
		min := -1
		for i := 0; i < len(lists); i++ {
			if nil == lists[i] {
				if i == len(lists) - 1 {
					lists = lists[:i]
				} else {
					lists = append(lists[:i], lists[i + 1:]...)
				}
				i--
				continue
			}
			if min == -1 || lists[i].Val < lists[min].Val {
				min = i
			}
		}
		if min == -1 {
			break;
		}
		tail.Next = lists[min]
		tail = tail.Next
		lists[min] = lists[min].Next
	}
	

	return head.Next
}