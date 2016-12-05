package demo

import "testing"

func Test_isUgly(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"normal", args{0}, false},
		{"normal", args{1}, true},
		{"normal", args{2}, true},
		{"normal", args{3}, true},
		{"normal", args{4}, true},
		{"normal", args{5}, true},
		{"normal", args{6}, true},
		{"normal", args{7}, false},
		{"normal", args{8}, true},
		{"normal", args{9}, true},
		{"normal", args{10}, true},
		{"normal", args{11}, false},
		{"normal", args{12}, true},
		{"normal", args{13}, false},
		{"normal", args{14}, false},
		{"normal", args{15}, true},
		{"normal", args{16}, true},
		{"normal", args{17}, false},
		{"normal", args{18}, true},
		{"normal", args{19}, false},
		{"normal", args{20}, true},
		{"normal", args{21}, false},
		{"normal", args{22}, false},
		{"normal", args{23}, false},
		{"normal", args{24}, true},
		{"normal", args{25}, true},
		{"normal", args{26}, false},
		{"normal", args{27}, true},
		{"normal", args{28}, false},
		{"normal", args{29}, false},
		{"normal", args{30}, true},
		{"normal", args{31}, false},
		{"normal", args{32}, true},
		{"normal", args{33}, false},
		{"normal", args{34}, false},
		{"normal", args{35}, false},
		{"normal", args{36}, true},
		{"normal", args{37}, false},
		{"normal", args{38}, false},
		{"normal", args{39}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUgly(tt.args.num); got != tt.want {
				t.Errorf("isUgly(%d) = %v, want %v", tt.args.num, got, tt.want)
			}
		})
	}
}
