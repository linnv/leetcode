package demo

import "sort"

func findDuplicates(nums []int) []int {
	sort.Ints(nums)
	ret := make([]int, 0, len(nums)/2)
	next := 0
	for i := 0; i < len(nums)-1; i++ {
		next = i + 1
		if nums[next] == nums[i] {
			ret = append(ret, nums[i])
			i++
		}

	}
	return ret
}
