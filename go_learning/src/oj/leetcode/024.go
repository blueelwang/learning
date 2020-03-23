package leetcode

func swapPairs(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return head
	}
	var last *ListNode = nil
	var newHead *ListNode = head.Next
	var p1 *ListNode = head
	var p2 *ListNode = head.Next

	for {
		if nil != last {
			last.Next = p2
		}
		p1.Next = p2.Next
		p2.Next = p1

		if nil == p1.Next || nil == p1.Next.Next {
			break
		}
		last = p1
		p1 = p1.Next
		p2 = p1.Next
	}
	
	return newHead
}