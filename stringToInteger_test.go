package demo

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"edge", args{"a"}, 0},
		{"edge", args{"1a"}, 1},
		{"edge", args{"+a1a"}, 0},
		{"edge", args{"+2."}, 2},
		{"edge", args{".+2."}, 0},
		{"edge", args{"--++-+++.1."}, 0},

		{"normal", args{"0"}, 0},
		{"normal", args{"-0"}, 0},
		{"normal", args{"-1"}, -1},
		{"normal", args{"+1"}, 1},
		{"normal", args{"+1."}, 1},
		{"normal", args{"-1."}, -1},
		{"normal", args{"00000--++-+++.1."}, 0},
		{"normal", args{"0001.0--++-+++.1."}, 1},

		{"edge", args{"-2147483648."}, -2147483648},
		{"edge", args{"2147483647."}, 2147483647},
		{"edge", args{"2147483648."}, 2147483647},
		{"normal", args{"2147483648.+2147483648.2147483648."}, 2147483647},
		{"normal", args{"-2147483649."}, -2147483648},
		{"normal", args{"-21474836.49"}, -21474836},
		{"normal", args{"+21474836.49"}, 21474836},
		{"normal", args{"+2.1474836.49"}, 2},
		{"normal", args{"-2.1474836.49"}, -2},

		{"normal", args{"+11191657170"}, 2147483647},
		{"normal", args{"+11191657170"}, 2147483647},
		{"normal", args{"18446744073709551617"}, 2147483647},
		{"normal", args{"-2147483647"}, -2147483647},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
