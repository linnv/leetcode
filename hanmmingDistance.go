package demo

func hammingDistance(x int, y int) int {
	z := x ^ y
	count := 0
	for z > 0 {
		count++
		z = z & (z - 1)
	}
	return count
}
