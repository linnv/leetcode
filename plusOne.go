package demo

func plusOne(digits []int) []int {
	digitsLen := len(digits) - 1
	carry := 0
	carry = (digits[digitsLen] + 1) / 10
	digits[digitsLen] = (digits[digitsLen] + 1) % 10
	for i := digitsLen - 1; i >= 0; i-- {
		if carry > 0 {
			carry = (digits[i] + 1) / 10
			digits[i] = (digits[i] + 1) % 10
		}
	}
	if carry > 0 {
		ret := make([]int, digitsLen+2)
		ret[0] = carry
		copy(ret[1:], digits)
		return ret
	}
	return digits
}
