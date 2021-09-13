package rollback

/*
1 2 3 4
递归 以某个数为开头/结尾时（swap操作） 剩下的数的排列
排列是指从n（总数）中有顺序的去除m个数排
组合是指从n（总数）中取出m个数排 不要求有序	m个做全排列 所以公式为(n!)/(n-m)!m!
全排列 排列中 m = n 的情况
*/

// 1、路径：也就是已经做出的选择。
//
// 2、选择列表：也就是你当前可以做的选择。
//
// 3、结束条件：也就是到达决策树底层，无法再做选择的条件。

// 其核心就是 for 循环里面的递归，在递归调用之前「做选择」，在递归调用之后「撤销选择」

// 全排列
var ret = make([][]int, 0)

func rcl(arr []int) [][]int {
	frcl(arr, 0, len(arr))
	return ret
}

func frcl(arr []int, ind, lth int) {
	if ind == lth {
		ret = append(ret, copyArr(arr))
		return
	}
	for i := ind; i < lth; i++ {
		swap(arr, i, ind)
		frcl(arr, ind+1, lth)
		swap(arr, i, ind)
	}
}

func copyArr(arr []int) []int {
	ret := make([]int, len(arr))
	for i, v := range arr {
		ret[i] = v
	}
	return ret
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func Nbacktrack() {
	// ...
}

func xxx() {
	// 3*3的组合 9宫格的遍历 双for循环
	// 9宫格的判断 某些点  for 双变量   + - 1
}
