package demo

// if positive numbers only, sum them all
// if negative numbers only, select the biggest number
// if positive and negative numbers exist both, sum all positive and plus the biggest negative
func maxSubArray(nums []int) int {
	numsLen := len(nums)
	if numsLen < 1 {
		return 0
	}
	previouse, max := nums[0], nums[0]
	for i := 1; i < numsLen; i++ {
		if previouse > 0 {
			previouse += nums[i]
		} else {
			previouse = nums[i]
		}
		if previouse > max {
			max = previouse
		}
	}
	return max
}
