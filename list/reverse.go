package list

import "fmt"

type Node struct {
	Val  string
	Next *Node
}

// 翻转链表
func reverse(head *Node) *Node {
	var pre, cur, next *Node = nil, head, nil
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func middleNode(head *Node) *Node {
	var slow, fast *Node = head, head
	for fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	if fast == nil {
		fmt.Println("双数节点数")
	} else {
		fmt.Println("单数节点数")
		return slow
	}
	return slow
}
