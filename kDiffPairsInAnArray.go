package demo

import "strconv"

func findPairs(nums []int, k int) int {
	numsLen := len(nums)
	if numsLen < 2 || k < 0 {
		return 0
	}
	pairCount := make(map[string]struct{}, numsLen)
	for i := 0; i < numsLen; i++ {
		for j := i + 1; j < numsLen; j++ {
			if (nums[i]-nums[j]) == k || ((nums[j] - nums[i]) == k) {
				if nums[i] > nums[j] {
					pairCount[strconv.Itoa(nums[j])+"_"+strconv.Itoa(nums[i])] = struct{}{}
					continue
				}
				pairCount[strconv.Itoa(nums[i])+"_"+strconv.Itoa(nums[j])] = struct{}{}
			}
		}

	}
	return len(pairCount)
}
