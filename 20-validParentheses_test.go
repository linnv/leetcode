package demo

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{

		{"edge", args{"["}, false},
		{"edge", args{"[["}, false},
		{"normal", args{"([])"}, true},
		{"edge", args{"([]]"}, false},
		{"edge", args{"([)]"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid(%s) = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
}
