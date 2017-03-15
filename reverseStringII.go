package demo

func reverseStrII(s string, k int) string {
	reverseAll := func(s string) string {
		byteS := []rune(s)
		sLen := len(byteS)
		for i := 0; i < sLen/2; i++ {
			byteS[i], byteS[sLen-i-1] = byteS[sLen-i-1], byteS[i]
		}
		return string(byteS)
	}

	sLen := len(s)
	if sLen <= k {
		return reverseAll(s)
	}
	reversePart := reverseAll(s[:k])
	if sLen <= 2*k {
		return reversePart + s[k:]
	}
	ret := ""
	startIndex, i := 0, 0
	for i = 0; i < sLen/(2*k); i++ {
		startIndex = i * 2 * k
		ret += reverseAll(s[startIndex:(startIndex+k)]) + s[(startIndex+k):(startIndex+2*k)]
	}
	startIndex = i * 2 * k
	if (sLen - startIndex) <= k {
		ret += reverseAll(s[startIndex:])
		return ret
	}
	ret += reverseAll(s[startIndex:(startIndex+k)]) + s[startIndex+k:]
	return ret
}
