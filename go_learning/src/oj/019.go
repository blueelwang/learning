package oj


type ListNode struct {
    Val int
    Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if nil == head || 0 == n {
		return head
	}

	len := 0
	for p := head; nil != p; p = p.Next {
		len++
		
	}
	
	n = len - n + 1
	if (n <= 0) {
		return head
	}
	var last *ListNode = nil
	for p, i := head, 1; nil != p; p, last, i = p.Next, p, i + 1 {
		if i == n {
			if nil == last {
				return p.Next
			} else {
				last.Next = p.Next
				break
			}
		}
	}
	return head
}