package demo

func isPerfectSquare(num int) bool {
	if num < 1 {
		return false
	}
	left, mid, right := 0, 0, num
	for left <= right {
		mid = (left + right) / 2
		t := mid * mid
		if t == num {
			return true
		}
		if t > num {
			right = mid - 1
		}
		if t < num {
			left = mid + 1
		}
	}
	return false
}
