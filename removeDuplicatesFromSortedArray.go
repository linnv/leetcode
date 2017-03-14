package demo

//note: testcase is not enough to test nums
func removeDuplicates(nums []int) int {
	numsLen := len(nums)
	if numsLen < 1 {
		return 0
	}
	pre := nums[0]
	count := 1
	newIndex := 1
	for i := 1; i < numsLen; i++ {
		if pre != nums[i] {
			pre = nums[i]
			nums[newIndex] = pre
			count++
			newIndex++
		}
	}
	nums = nums[:newIndex]
	return count
}
