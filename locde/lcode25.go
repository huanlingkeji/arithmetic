package locde

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	a, b := head, head
	for i := 0; i < k; i++ {
		for b == nil {
			return head
		}
		b = b.Next
	}
	var newHead = reverse(a, b)
	a.Next = reverseKGroup(b, k)
	return newHead
}

func reverse(a, b *ListNode) *ListNode {
	var pre, cur, next *ListNode = nil, a, a
	for cur != b {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
