package locde

import "math"

func nodesBetweenCriticalPoints(head *ListNode) []int {
	left, pre, cur := -1, -1, -1
	min := math.MaxInt
	var preN, curN, nextN *ListNode = nil, head, head
	no := 0
	for curN != nil {
		nextN = curN.Next
		if preN != nil && nextN != nil {
			if (preN.Val > curN.Val && nextN.Val > curN.Val) || (preN.Val < curN.Val && nextN.Val < curN.Val) {
				if left < 0 {
					left = no
				}
				if no-left > 0 && min > no-pre {
					min = no - pre
				}
				pre = no
				cur = no
			}
		}
		preN = curN
		curN = curN.Next
		no++
	}
	if min != 100000 && left > 0 {
		return []int{min, cur - left}
	}
	return []int{-1, -1}
}
