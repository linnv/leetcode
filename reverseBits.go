package demo

func reverseBits(n int) int {
	ret := 0
	for i := 0; i < 32; i++ {
		ret = (ret << 1) | (n & 1) //it doesn't matter the first time shiting zero to left
		n >>= 1
	}
	return ret
}
