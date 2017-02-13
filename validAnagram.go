package demo

// better challenge for next is 438
// space for time
func isAnagram(s string, t string) bool {
	sLen := len(s)
	tLen := len(t)
	if sLen != tLen {
		return false
	}
	exist := make([]int, int('z')+1)
	for i := 0; i < tLen; i++ {
		exist[int(t[i])]++
	}
	for i := 0; i < sLen; i++ {
		exist[int(s[i])]--
		if exist[int(s[i])] < 0 {
			return false
		}
	}
	return true
}
