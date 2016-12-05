package demo

// found no direct rule to solve this problem but as obviously as the information has provided
func isHappy(n int) bool {
	unHappyNum := make([]bool, 11)
	unHappyNum[1], unHappyNum[7], unHappyNum[10] = true, true, true

	if n < 10 {
		return unHappyNum[n]
	}
	processedNum := make(map[int]bool)
	processedNum[n] = true
	for {
		if n < 11 {
			return unHappyNum[n]

		}
		tmp := n
		n = 0
		for tmp > 0 {
			n += (tmp % 10) * (tmp % 10)
			tmp /= 10
		}
		tmp = n
		if processedNum[n] {
			return false
		}

	}
}
