package sort

import (
	"arithmetic/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkQuickSort(t *testing.B) {
	testNum := 10
	arr := rand.BuildRandomArr(testNum)
	QuickSort(arr, 0, testNum-1)
	for i := 0; i < testNum; i++ {
		assert.Equal(t, i, arr[i])
	}
}

func TestName(t *testing.T) {
	testNum := 10
	arr := rand.BuildRandomArr(testNum)
	arr = []int{1, 7, 9, 3, 5, 8, 0, 2, 4, 6}
	QuickSort(arr, 0, testNum-1)
	for i := 0; i < testNum; i++ {
		assert.Equal(t, i, arr[i])
	}
}
