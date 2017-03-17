package demo

import "sort"

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	numsLen := len(nums)
	if numsLen < 1 {
		return false
	}
	for i := 1; i < numsLen; i++ {
		if nums[i-1] == nums[i] {
			return true
		}
	}
	return false
}
