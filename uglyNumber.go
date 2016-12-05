package demo

func divideNum(num int) bool {
	if num < 6 {
		return true
	}
	if num%2 == 0 {
		return divideNum(num / 2)
	}
	if num%3 == 0 {
		return divideNum(num / 3)
	}
	if num%5 == 0 {
		return divideNum(num / 5)
	}
	return false
}

func isUgly(num int) bool {
	if num < 1 {
		return false
	}
	return divideNum(num)
}
