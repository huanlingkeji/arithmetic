package sort

import (
	"fmt"
	"testing"
)

//如果从大到小 则移动到最后 假设小找大  则会找不到 i的值就会移动到最后
//如果从小到大 则不用移动
//i（找比之大的值） j（找比之小的值）
//

//因为是要比比之小的值往左边挪  所以应该是先找比之小的值

//快排序算法  对于未排序完成的数组  选择一个基数（可默认为第一个） 然后循环选取一个比之小的数和比之大的数进行交换 最后把该数字移动到中 则数组左边都比它小 后边都比它大
func TestQuitSort(t *testing.T) {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	quicksort(arr, 0, len(arr)-1)
	fmt.Println("arr: ", arr)
}

func quicksort(arr []int, l, r int) {
	if l >= r {
		return
	}
	pIndx := getposition(arr, l, r)
	quicksort(arr, l, pIndx-1)
	quicksort(arr, pIndx+1, r)
}

func getposition(arr []int, l, r int) int {
	i, j := l, r
	temp := arr[i]
	for i < j {
		for arr[j] >= temp && i < j {
			j--
		}
		for arr[i] <= temp && i < j {
			i++
		}
		swap(arr, i, j)
	}
	swap(arr, l, i)
	return i
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
