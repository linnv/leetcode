package demo

func strStr(haystack string, needle string) int {
	haystackLen := len(haystack)
	needleLen := len(needle)
	found := true
	for i := 0; i < haystackLen-needleLen+1; i++ {
		found = true
		for j := 0; j < needleLen; j++ {
			if haystack[i+j] != needle[j] {
				found = false
				break
			}
		}
		if found {
			return i
		}
	}
	return -1
}
