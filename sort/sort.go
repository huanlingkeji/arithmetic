package sort

import "arithmetic/util"

/*
快排:选定一个基准值(对数组以此基准进行排序，将比它小的放一边，比它大的放另一边).对左右两边递归进行此步骤。
*/

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func doQuickSort(arr []int, i, j int) int {
	temp := arr[i]
	l := i
	r := j
	for l < r {
		// 找到第一个比基准值小的值
		for arr[r] >= temp && l < r {
			r--
		}
		// 找到第一个比基准值大的值
		for arr[l] <= temp && l < r {
			l++
		}
		// 注意：先找比基准打的值会造成，当没有交换时 l的位置不正确
		swap(arr, l, r)
	}
	swap(arr, i, l)
	return l
}

func QuickSort(arr []int, i, j int) {
	util.PrintArr(arr)
	if i < j {
		ind := doQuickSort(arr, i, j)
		QuickSort(arr, i, ind-1)
		QuickSort(arr, ind+1, j)
	}
}
