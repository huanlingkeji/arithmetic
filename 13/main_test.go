package _3

import (
	"fmt"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		key int
		val AnyType
	}{
		{
			key: 1,
			val: 1,
		},
		{
			key: 2,
			val: 2,
		},
		{
			key: 3,
			val: 3,
		},
		{
			key: 2,
			val: 2,
		},
		{
			key: 3,
			val: 3,
		},
		{
			key: 1,
			val: 1,
		},
		{
			key: 2,
			val: 2,
		},
	}
	lru := NewLRU(3)
	for _, tt := range tests {
		lru.Set(tt.key, tt.val)
		for i := lru.list.Back(); i != nil; i = i.Prev() {
			fmt.Print(i.Value.(pair).Val, " ")
		}
		fmt.Println()
	}
}
