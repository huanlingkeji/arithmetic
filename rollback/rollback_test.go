package rollback

import (
	"fmt"
	"testing"
)

func Test_rcl(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test1",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rcl(tt.args.arr)
			fmt.Println("got", got)
		})
	}
}
