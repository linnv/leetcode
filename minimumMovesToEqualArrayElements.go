package demo

import "sort"

func minMoves(nums []int) int {
	sort.Ints(nums)
	numsLen := len(nums)
	sum := 0
	for i := 0; i < numsLen; i++ {
		sum += nums[i]
	}
	return sum - nums[0]*numsLen
}
