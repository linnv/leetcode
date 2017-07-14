package demo

func jump(num []int) int {
	numLen := len(num)
	if numLen < 2 {
		return 0
	}
	count := 0
	jumpIndex, jumpNum, maxJumpWeight, jumpEnd := 0, 0, 0, 0
	for {
		leftJump := numLen - jumpIndex
		if leftJump == 0 {
			return count
		}
		jumpNum = num[jumpIndex]
		if jumpNum > leftJump {
			count++
			return count
		}

		maxJumpWeight = 0
		jumpEnd = jumpNum + jumpIndex
		for j := jumpIndex + 1; j <= jumpEnd; j++ {
			if j == numLen-1 {
				count++
				return count
			}
			if maxJumpWeight <= num[j]+j {
				maxJumpWeight = num[j] + j
				jumpIndex = j
			}
		}
		count++
	}
	return count
}
