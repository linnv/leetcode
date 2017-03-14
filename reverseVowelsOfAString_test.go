package demo

import "testing"

func Test_reverseVowel(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"normal", args{"hello"}, "holle"},
		{"normal", args{"leetcode"}, "leotcede"},
		{"normal", args{"eo"}, "oe"},
		{"normal", args{"Sore was I ere I saw Eros."}, "SorE was I ere I saw eros."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseVowel(tt.args.s); got != tt.want {
				t.Errorf("reverseVowel() = %v, want %v", got, tt.want)
			}
		})
	}
}
