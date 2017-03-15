package demo

func findComplement(num int) int {
	mask := 1
	n := num
	for n > 0 {
		mask <<= 1
		n >>= 1
	}
	return num ^ (mask - 1)

}
