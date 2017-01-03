package demo

//space for time
func findDisappearedNumbers(nums []int) []int {
	numsLen := len(nums)
	numsMark := make([]bool, numsLen+1)
	for i := 0; i < numsLen; i++ {
		numsMark[nums[i]] = true
	}
	ret := make([]int, 0, numsLen/2)
	for i := 1; i <= numsLen; i++ {
		if !numsMark[i] {
			ret = append(ret, i)
		}
	}
	return ret
}
