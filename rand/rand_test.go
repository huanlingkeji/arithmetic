package rand

import (
	"reflect"
	"testing"
)

func TestGetRandomArr(t *testing.T) {
	type args struct {
		arrLen int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test",
			args: args{
				arrLen: 10,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildRandomArr(tt.args.arrLen); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildRandomArr() = %v, want %v", got, tt.want)
			}
		})
	}
}
