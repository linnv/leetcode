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
}

func jump2(num []int) int {
	numLen := len(num)
	target := numLen - 1
	if target == 0 {
		return 0
	}
	var maxCanreach, reach, count int
	for i := 0; i < numLen; i++ {
		canReach := num[i] + i
		if canReach >= target {
			break
		}
		if reach < canReach {
			reach = canReach
		}

		if i >= maxCanreach {
			maxCanreach = reach
			count++
		}
	}
	count++
	return count
}
