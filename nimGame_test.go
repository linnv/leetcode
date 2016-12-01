package demo

import "testing"

func Test_canWinNim(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"edge", args{1}, true},
		{"edge", args{2}, true},
		{"edge", args{3}, true},
		{"edge", args{4}, false},
		{"edge", args{5}, true},
		{"edge", args{6}, true},
		{"edge", args{7}, true},
		{"edge", args{8}, false},
		{"edge", args{9}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canWinNim(tt.args.n); got != tt.want {
				t.Errorf("canWinNim() = %v, want %v", got, tt.want)
			}
		})
	}
}
