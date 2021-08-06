package str

import "testing"

func Test_ns(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				s: "UUUUU",
			},
			want: "uuuuu",
		},
		{
			name: "test2",
			args: args{
				s: "userID",
			},
			want: "user_id",
		},
		{
			name: "test3",
			args: args{
				s: "aaaBbbCCC",
			},
			want: "aaa_bbb_ccc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ns(tt.args.s); got != tt.want {
				t.Errorf("ns() = %v, want %v", got, tt.want)
			}
		})
	}
}
