package demo

import "testing"

func Test_reverseBits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"normal", args{1}, 1 << 31},
		{"normal", args{2}, 1 << 30},
		{"normal", args{43261596}, 964176192},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBits(tt.args.n); got != tt.want {
				t.Errorf("reverseBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
