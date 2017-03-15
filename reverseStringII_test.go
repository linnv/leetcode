package demo

import "testing"

func Test_reverseStrII(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"normal", args{"abcdefg", 2}, "bacdfeg"},
		{"normal", args{"abc", 2}, "bac"},
		{"normal", args{"ababcb", 2}, "baabbc"},
		{"normal", args{"abc", 3}, "cba"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseStrII(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("reverseStrII() = %v, want %v", got, tt.want)
			}
		})
	}
}
