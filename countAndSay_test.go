package demo

import "testing"

func Test_countAndSay(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"normal", args{1}, "1"},
		{"normal", args{2}, "11"},
		{"normal", args{3}, "21"},
		{"normal", args{4}, "1211"},
		{"normal", args{5}, "111221"},
		{"normal", args{6}, "312211"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countAndSay(tt.args.n); got != tt.want {
				t.Errorf("countAndSay() = %v, want %v", got, tt.want)
			}
		})
	}
}
