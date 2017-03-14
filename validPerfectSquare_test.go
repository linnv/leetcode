package demo

import "testing"

func Test_isPerfectSquare(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"normal", args{4}, true},
		{"normal", args{1}, true},
		{"normal", args{9}, true},
		{"normal", args{16}, true},
		{"normal", args{7}, false},
		{"normal", args{15}, false},
		{"normal", args{8}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPerfectSquare(tt.args.num); got != tt.want {
				t.Errorf("isPerfectSquare(%d) = %v, want %v", tt.args.num, got, tt.want)
			}
		})
	}
}
