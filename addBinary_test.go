package demo

import "testing"

func Test_addBinanry(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"edge", args{"1010", "1011"}, "10101"},
		{"normal", args{"11", "1"}, "100"},
		{"edge", args{"11", "11"}, "110"},
		{"edge", args{"1", "0"}, "1"},
		{"edge", args{"0", "0"}, "0"},
		{"edge", args{"0", ""}, "0"},
		{"edge", args{"1", ""}, "1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addBinanry(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("addBinanry(%s,%s) = %v, want %v", tt.args.a, tt.args.b, got, tt.want)
			}
		})
	}
}
