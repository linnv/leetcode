package demo

import "strings"

func countSegments(s string) int {
	s = strings.TrimSpace(s)
	sLen := len(s)
	if sLen < 1 {
		return 0
	}
	count := 0
	spaceInt := int(byte(' '))
	for i := 0; i < sLen-1; i++ {
		if int(s[i]) != spaceInt && int(s[i+1]) == spaceInt {
			count++
		}

	}
	return count + 1
}
