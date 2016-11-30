package demo

import (
	"reflect"
	"testing"
)

func Test_fizzBuzz(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"edge", args{n: 2}, []string{"1", "2"}},
		{"edge", args{n: 3}, []string{"1", "2", "Fizz"}},
		{"edge", args{5}, []string{"1", "2", "Fizz", "4", "Buzz"}},
		{"normal", args{16}, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz", "16"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fizzBuzz(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}
