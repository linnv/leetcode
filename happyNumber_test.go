package demo

import "testing"

func Test_isHappy(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"normal", args{0}, false},
		{"normal", args{1}, true},
		{"normal", args{2}, false},
		{"normal", args{3}, false},
		{"normal", args{4}, false},
		{"normal", args{5}, false},
		{"normal", args{6}, false},
		{"normal", args{7}, true},
		{"normal", args{8}, false},
		{"normal", args{9}, false},
		{"normal", args{10}, true},
		{"normal", args{11}, false},
		{"normal", args{12}, false},
		{"edge", args{13}, true},
		{"normal", args{14}, false},
		{"normal", args{15}, false},
		{"normal", args{16}, false},
		{"normal", args{17}, false},
		{"normal", args{18}, false},
		{"normal", args{19}, true},
		{"normal", args{20}, false},
		{"normal", args{21}, false},
		{"normal", args{22}, false},
		{"normal", args{23}, true},
		{"normal", args{24}, false},
		{"normal", args{25}, false},
		{"normal", args{26}, false},
		{"normal", args{27}, false},
		{"normal", args{28}, true},
		{"normal", args{29}, false},
		{"normal", args{30}, false},
		{"normal", args{31}, true},
		{"normal", args{32}, true},
		{"normal", args{33}, false},
		{"normal", args{34}, false},
		{"normal", args{35}, false},
		{"normal", args{36}, false},
		{"normal", args{37}, false},
		{"normal", args{38}, false},
		{"normal", args{39}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHappy(tt.args.n); got != tt.want {
				t.Errorf("isHappy(%d) = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
}
