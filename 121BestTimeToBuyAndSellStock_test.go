package demo

import (
	"testing"
)

func TestBestTimeToBuyAndSellStock(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"normal", args{[]int{7, 1, 5, 3, 6, 4}}, 5},
		{"edge", args{[]int{1, 1, 0}}, 0},
		{"normal", args{[]int{1}}, 0},
		{"normal", args{[]int{}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BestTimeToBuyAndSellStock(tt.args.nums); got != tt.want {
				t.Errorf("BestTimeToBuyAndSellStock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBestTimeToBuyAndSellStockViolence(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"normal", args{[]int{7, 1, 5, 3, 6, 4}}, 5},
		{"edge", args{[]int{1, 1, 0}}, 0},
		{"normal", args{[]int{1}}, 0},
		{"normal", args{[]int{}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BestTimeToBuyAndSellStockViolence(tt.args.nums); got != tt.want {
				t.Errorf("BestTimeToBuyAndSellStockViolence() = %v, want %v", got, tt.want)
			}
		})
	}
}
