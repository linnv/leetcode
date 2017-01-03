package demo

import "sort"

func firstMissingPositive(nums []int) int {
	if len(nums) < 1 {
		return 1
	}
	sort.Ints(nums)
	next := 0
	for i := 0; i < len(nums)-1; i++ {
		next = i + 1
		if nums[i] >= 0 {
			if i == 0 && nums[i] > 1 {
				return 1
			}
			if nums[next] == nums[i] {
				continue
			}
			if nums[next] != nums[i]+1 {
				return nums[i] + 1
			}
			continue
		}
		if nums[next] == nums[i] {
			continue
		}
		if nums[next] > 0 && nums[next] != 1 {
			return 1
		}
	}
	if nums[len(nums)-1] <= 0 || (len(nums) == 1 && nums[len(nums)-1] > 1) {
		return 1
	}

	return nums[len(nums)-1] + 1
}
