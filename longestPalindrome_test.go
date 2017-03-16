package demo

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"normal", args{"abccccdd"}, 7},
		{"normal", args{"aa"}, 2},
		{"normal", args{"aaa"}, 3},
		{"normal", args{"zz"}, 2},
		{"normal", args{"zzz"}, 3},
		{"normal", args{"ZZ"}, 2},
		{"normal", args{"aba"}, 3},
		{"normal", args{"ABA"}, 3},
		{"normal", args{"a"}, 1},
		{"normal", args{"A"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
