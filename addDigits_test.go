package demo

import "testing"

func Test_addDigits(t *testing.T) {
	type args struct {
		nums int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"normal", args{0}, 0},
		{"normal", args{1}, 1},
		{"normal", args{2}, 2},
		{"normal", args{3}, 3},
		{"normal", args{4}, 4},
		{"normal", args{5}, 5},
		{"normal", args{6}, 6},
		{"normal", args{7}, 7},
		{"normal", args{8}, 8},
		{"normal", args{9}, 9},
		{"edge", args{10}, 1},
		{"normal", args{11}, 2},
		{"normal", args{12}, 3},
		{"normal", args{13}, 4},
		{"normal", args{14}, 5},
		{"normal", args{15}, 6},
		{"normal", args{16}, 7},
		{"normal", args{17}, 8},
		{"edge", args{18}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addDigits(tt.args.nums); got != tt.want {
				t.Errorf("addDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
