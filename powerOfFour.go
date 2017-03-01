package demo

func isPowerOfFour(num int) bool {
	isPowerOfTwo := (num > 0) && (num&(num-1) == 0)
	if !isPowerOfTwo {
		return false
	}
	//0x555555
	//0101=5,0101 x 8=32bit,actuall 31bit is enough
	return isPowerOfTwo && ((num & 0x55555555) == num)
}
