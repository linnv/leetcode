package demo

func singleNumber2(nums []int) int {
	bitsCount := make([]int, 32)
	var result int
	var i uint
	var minimalCount int
	// int in x64 is 64 bits long, so we handle this minimal number while leetcode treats int to always be 32 bits long
	// var i int
	// var i64 int64
	// i = 1
	// i64 = 1
	// log.Printf("unsafe.Sizeof(i): %+v\n", unsafe.Sizeof(i))
	// log.Printf("unsafe.Sizeof(i64): %+v\n", unsafe.Sizeof(i64))
	// log.Printf("i<<: %+v\n", i<<(unsafe.Sizeof(i)*8-1))
	// log.Printf("i64<<: %+v\n", i64<<(unsafe.Sizeof(i64)*8-1))
	// log.Printf("i<<: %+v\n", i<<(unsafe.Sizeof(i)*8-1))
	const miniNum = 2147483648
	for j := 0; j < len(nums); j++ {
		unsigneNumer := nums[j]
		if unsigneNumer < 0 {
			bitsCount[31]++
			unsigneNumer = -unsigneNumer
		}
		for i = 0; i < 31; i++ {
			if unsigneNumer == miniNum {
				minimalCount++
				continue
			}
			if ((unsigneNumer >> i) & 1) == 1 {
				bitsCount[i]++
			}
		}
	}
	if minimalCount%3 != 0 {
		return -miniNum
	}
	for i = 0; i < 31; i++ {
		result |= ((bitsCount[i] % 3) << i)
	}
	if (bitsCount[31] % 3) > 0 {
		result = -result
	}
	return result
}
