package demo

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"normal", args{"bbacbababaabcbababbacbababaabcbababbacbababaabcbababbacbababaabcbababbacbababaabcbababbacbababaabcbababbacbababaabcbababacbababaabcbababababca", "abababca"}, 134},
		{"normal", args{"bac", "ac"}, 1},
		{"normal", args{"bacxxxx", "xx"}, 3},
		{"normal", args{"bcac", "c"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
