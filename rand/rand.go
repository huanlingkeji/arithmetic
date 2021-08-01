package rand

import "math/rand"

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func less(arr []int, i, j int) bool {
	return arr[i] < arr[j]
}

func length(arr []int) int {
	return len(arr)
}

func BuildRandomArr(arrLen int) []int {
	retArr := make([]int, arrLen)
	for i := 0; i < arrLen; i++ {
		retArr[i] = i
	}
	for i := 0; i < arrLen; i++ {
		j := rand.Intn(arrLen-i) + i
		swap(retArr, i, j)
	}
	return retArr
}
