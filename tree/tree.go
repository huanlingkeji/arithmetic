package tree

import "fmt"

// getIndex获取的是元素的坐标 idx
// 如果是想使用该元素前的所有元素 则为arr[:idx]
// 如果要包括该元素则为arr[:idx+1]

type ListNode struct {
	Val   int
	Left  *ListNode
	Right *ListNode
}

func (t *ListNode) PreTraverse() {
	if t == nil {
		return
	}
	fmt.Println(t.Val, " ")
	t.Left.PreTraverse()
	t.Right.PreTraverse()
}

func (t *ListNode) MinTraverse() {
	if t == nil {
		return
	}
	t.Left.MinTraverse()
	fmt.Println(t.Val, " ")
	t.Right.MinTraverse()
}

func (t *ListNode) PostorderTraverse() {
	if t == nil {
		return
	}
	t.Left.PostorderTraverse()
	t.Right.PostorderTraverse()
	fmt.Println(t.Val, " ")
}

func (t *ListNode) LevelTraverse() {
	if t == nil {
		return
	}
	list := []*ListNode{t}
	for len(list) != 0 {
		node := list[0]
		fmt.Println(node.Val, " ")
		if node.Left != nil {
			list = append(list, node.Left)
		}
		if node.Right != nil {
			list = append(list, node.Right)
		}
		list = list[1:]
	}
}

func BuildTreeWithLevel(list []int, length, n int) *ListNode {
	var node *ListNode
	if n >= length {
		return nil
	}
	node = &ListNode{
		Val: list[n],
	}
	node.Left = BuildTreeWithLevel(list, length, 2*n+1)
	node.Right = BuildTreeWithLevel(list, length, 2*(n+1))
	return node
}

func BuildTreeWithMidPreTraverse(mid, pre []int) *ListNode {
	var node *ListNode
	if len(pre) == 0 {
		return nil
	}
	node = &ListNode{
		Val: pre[0],
	}
	idx := getIndexOfVal(mid, pre[0])
	node.Left = BuildTreeWithMidPreTraverse(mid[:idx], pre[1:idx+1])
	node.Right = BuildTreeWithMidPreTraverse(mid[idx+1:], pre[idx+1:])
	return node
}

func BuildTreeWithMidPostTraverse(mid, post []int) *ListNode {
	var node *ListNode
	if len(post) == 0 {
		return nil
	}
	postLen := len(post)
	val := post[postLen-1]
	node = &ListNode{
		Val: val,
	}
	idx := getIndexOfVal(mid, val)
	node.Left = BuildTreeWithMidPostTraverse(mid[:idx], post[:idx])
	node.Right = BuildTreeWithMidPostTraverse(mid[idx+1:], post[idx:postLen-1])
	return node
}

func getIndexOfVal(list []int, val int) int {
	for i, v := range list {
		if v == val {
			return i
		}
	}
	return -1
}

func GetTestTree() *ListNode {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	return BuildTreeWithLevel(list, len(list), 0)
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
