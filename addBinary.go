package demo

func addBinanry(a string, b string) string {
	minLen := len(a)
	maxLen := len(b)
	if minLen == 0 || maxLen == 0 {
		return a + b
	}
	var minSlice, maxSlice []byte
	if minLen > maxLen {
		minLen, maxLen = maxLen, minLen
		maxSlice, minSlice = []byte(a), []byte(b)
	} else {
		maxSlice, minSlice = []byte(b), []byte(a)
	}

	oneByte, zeroByte := byte('1'), byte('0')
	calBit := func(a, b byte) (byte, byte) {
		if (int(a) & int(b)) == int(oneByte) {
			return oneByte, zeroByte
		}
		if (int(a) | int(b)) == int(oneByte) {
			return zeroByte, oneByte
		}
		return zeroByte, zeroByte
	}
	carry := zeroByte
	bit := zeroByte //bit and carry should put outside for loop
	carry, maxSlice[maxLen-1] = calBit(minSlice[minLen-1], maxSlice[maxLen-1])
	for i := minLen - 2; i >= 0; i-- {
		carry, bit = calBit(minSlice[i], carry)
		if carry == oneByte {
			continue
		}
		carry, maxSlice[maxLen-minLen+i] = calBit(bit, maxSlice[maxLen-minLen+i])
	}
	if carry == oneByte {
		if minLen != maxLen {
			for i := maxLen - minLen - 1; i >= 0; i-- {
				carry, maxSlice[i] = calBit(carry, maxSlice[i])
			}
			if carry == oneByte {
				goto newSlice
			}
			return string(maxSlice)
		}
	newSlice:
		ret := make([]byte, maxLen+1)
		ret[0] = oneByte
		copy(ret[1:], maxSlice)
		return string(ret)
	}
	return string(maxSlice)
}
