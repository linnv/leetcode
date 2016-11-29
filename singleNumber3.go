package demo

func singleNumber3(nums []int) []int {
	var oneBitMark, zeroBitMark, oneOrZero int
	for i := 0; i < len(nums); i++ {
		oneOrZero ^= nums[i]
	}
	// we find the bit that can reconize these two number
	mark := 1
	for (oneOrZero & 1) != 1 {
		oneOrZero >>= 1
		mark <<= 1
	}

	for i := 0; i < len(nums); i++ {
		if nums[i]&mark > 0 {
			oneBitMark ^= nums[i]
		} else {
			zeroBitMark ^= nums[i]
		}
	}

	return []int{oneBitMark, zeroBitMark}
}
