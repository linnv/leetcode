package demo

import "testing"

func Test_islandPerimeter(t *testing.T) {
	arrayOne := [][]int{}
	arrayOne = append(arrayOne, []int{0, 1, 0, 0})
	arrayOne = append(arrayOne, []int{1, 1, 1, 0})
	arrayOne = append(arrayOne, []int{0, 1, 0, 0})
	arrayOne = append(arrayOne, []int{1, 1, 0, 0})

	arrayTwo := [][]int{}
	arrayTwo = append(arrayTwo, []int{0, 1, 0, 0})
	arrayTwo = append(arrayTwo, []int{1, 1, 1, 0})

	arrayThree := [][]int{}
	arrayThree = append(arrayThree, []int{1, 1})

	tests := []struct {
		name string
		args [][]int
		want int
	}{
		{"normal", arrayOne, 16},
		{"normal", arrayTwo, 10},
		{"edge", arrayThree, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := islandPerimeter(tt.args); got != tt.want {
				t.Errorf("islandPerimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}
