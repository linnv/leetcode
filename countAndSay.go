package demo

import "strconv"

func countAndSay(n int) string {
	newBirt := func(s string) string {
		var ret string
		for i := 0; i < len(s); {
			count := 0
			if i == len(s)-1 {
				ret = ret + strconv.Itoa(1) + string(s[i])
				return ret
			}
			for j := i + 1; j < len(s); j++ {
				count++
				if s[i] != s[j] {
					ret = ret + strconv.Itoa(count) + string(s[i])
					i += count - 1
					break
				}
				if j == len(s)-1 {
					return ret + strconv.Itoa(count+1) + string(s[i])
				}
			}
			i++
		}
		return ret
	}

	previouse := "1"
	next := previouse
	for i := 0; i < n-1; i++ {
		next = newBirt(previouse)
		previouse = next
	}

	return next
}
