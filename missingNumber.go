package demo

func missingNumber(nums []int) int {
	numsLen := len(nums)
	total := numsLen * (numsLen + 1) / 2
	for i := 0; i < numsLen; i++ {
		total -= nums[i]
	}

	return total
}

func missingNumberByBitsOperation(nums []int) int {
	numsLen := len(nums)
	ret, i := 0, 0
	for i = 0; i < numsLen; i++ {
		ret ^= i ^ nums[i]
	}

	return ret ^ i
}

//space for time, not recommand
// func missingNumber(nums []int) int {
// 	numsLen := len(nums)
// 	numsMark := make([]bool, numsLen+1)
// 	for i := 0; i < numsLen; i++ {
// 		numsMark[nums[i]] = true
// 	}
// 	for i := 0; i < len(numsMark); i++ {
// 		if !numsMark[i] {
// 			return i
// 		}
// 	}
// 	return 0
// }
