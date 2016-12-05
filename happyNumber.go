package demo

// found no direct rule to solve this problem but as obviously as the information has provided
func isHappy(n int) bool {
	happyNum := make([]bool, 11)
	happyNum[1], happyNum[7], happyNum[10] = true, true, true

	if n < 11 {
		return happyNum[n]
	}
	processedNum := make(map[int]bool)
	processedNum[n] = true
	for {
		if n < 11 {
			return happyNum[n]

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
