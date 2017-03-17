package demo

func containsNearbyDuplicate(nums []int, k int) bool {
	numsLen := len(nums)
	if numsLen < 1 {
		return false
	}
	indexMap := make(map[int]int, numsLen)
	for i := 0; i < numsLen; i++ {
		if index, ok := indexMap[nums[i]]; ok && i-index <= k {
			return true
		}
		indexMap[nums[i]] = i
	}
	return false
}
