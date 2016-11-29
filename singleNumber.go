package demo

func singleNumber(nums []int) int {
	var j int
	for i := 0; i < len(nums); i++ {
		j ^= nums[i]
	}
	return j
}
