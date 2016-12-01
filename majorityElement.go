package demo

func majorityElement(nums []int) int {
	numsLen := len(nums)
	half := numsLen / 2
	count := 0
	majority := int(0)
	for i := 0; i < numsLen; i++ {
		// low efficient way
		// for j := i; j < numsLen; j++ {
		// 	if nums[i] == nums[j] {
		// 		count++
		// 		if count > half {
		// 			return nums[i]
		// 		}
		// 	}
		// }
		// count = 0

		// 1000+ms VS 40+ms

		//Moore Voting Algorithm
		//what we care is the counting of majority, count must be >=0
		if count == 0 {
			majority = nums[i]
			count++
			continue
		}
		if majority == nums[i] {
			count = count + 1
			if count > half {
				return majority
			}
			continue
		}
		count--
	}
	return majority
}
