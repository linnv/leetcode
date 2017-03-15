package demo

func findMaxConsecutiveOnes(nums []int) int {
	numsLen := len(nums)
	count := 0
	max := 0
	for i := 0; i < numsLen; i++ {
		if nums[i] == 1 {
			count++
			if count > max {
				max = count
			}
		} else {
			count = 0
		}
	}
	return max
}

//another way also work, but much code to write
// func findMaxConsecutiveOnes(nums []int) int {
// 	numsLen := len(nums)
// 	zero := make([]int, 0, numsLen)
// 	for i := 0; i < numsLen; i++ {
// 		if nums[i] != 1 {
// 			zero = append(zero, i)
// 		}
// 	}
// 	if nums[numsLen-1] == 1 {
// 		zero = append(zero, numsLen)
// 	}
// 	max := zero[0]
// 	zeroLen := len(zero)
// 	pre := zero[0]
// 	for i := 1; i < zeroLen; i++ {
// 		if zero[i]-pre-1 > max {
// 			max = zero[i] - pre - 1
// 		}
// 		pre = zero[i]
// 	}
// 	return max
// }
