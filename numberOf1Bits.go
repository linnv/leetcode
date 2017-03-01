package demo

func hanmmingWeight(n uint) int {
	count := 0
	for n > 0 {
		n = n & (n - 1)
		count++
	}
	return count
}
