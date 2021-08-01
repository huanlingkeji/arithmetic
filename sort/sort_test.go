package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"yys.com/v0/rand"
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
	QuickSort(arr, 0, testNum-1)
	for i := 0; i < testNum; i++ {
		assert.Equal(t, i, arr[i])
	}
}
