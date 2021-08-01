package tree

import (
	"fmt"
	"testing"
)

func TestBuildTree(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tree := BuildTreeWithLevel(list, len(list), 0)
	tree.LevelTraverse()
	tree.PostorderTraverse()
}

func TestBuildTreeWithMidPreTraverse(t *testing.T) {
	LevelArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // 层次
	PreArr := []int{1, 2, 4, 8, 9, 5, 10, 3, 6, 7}   // 先根
	MidArr := []int{8, 4, 9, 2, 5, 10, 1, 6, 3, 7}   // 中根
	PostArr := []int{8, 9, 4, 10, 5, 2, 6, 7, 3, 1}  // 后根
	tree := BuildTreeWithMidPreTraverse(MidArr, PreArr)
	tree.LevelTraverse()
	_ = LevelArr
	_ = PostArr
}

func TestBuildTreeWithMidPostTraverse(t *testing.T) {
	LevelArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // 层次
	PreArr := []int{1, 2, 4, 8, 9, 5, 10, 3, 6, 7}   // 先根
	MidArr := []int{8, 4, 9, 2, 5, 10, 1, 6, 3, 7}   // 中根
	PostArr := []int{8, 9, 4, 10, 5, 2, 6, 7, 3, 1}  // 后根
	tree := BuildTreeWithMidPostTraverse(MidArr, PostArr)
	tree.LevelTraverse()
	_ = LevelArr
	_ = PreArr
}

func TestLargeTreePath(t *testing.T) {
	testList := make([]int, 15)
	for i := range testList {
		testList[i] = i + 1
	}
	tree := BuildTreeWithLevel(testList, len(testList), 0)
	fmt.Println(getMaxTreePathVal(tree))
}

func getMaxTreePathVal(t *ListNode) int {
	if t == nil {
		return 0
	}
	return max(getMaxTreePathVal(t.Left), getMaxTreePathVal(t.Right)) + t.Val
}
