package demo

func findAnagrams(s string, p string) []int {
	pLen := len(p)
	sLen := len(s)
	ret := make([]int, 0, 10)
	if sLen == 0 || pLen > sLen {
		return ret
	}
	if pLen == 0 {
		for i := 0; i < sLen; i++ {
			ret = append(ret, i)
		}
		//what the hell is this
		ret = append(ret, sLen)
		return ret
	}

	constanLen := int('z') + 1
	subStr := make([]int, constanLen)
	processStr := make([]int, constanLen)
	for i := 0; i < pLen; i++ {
		subStr[int(p[i])]++
		processStr[int(s[i])]++
	}

	isAnagram := func(processStr, subStr []int) bool {
		for i := 0; i < constanLen; i++ {
			if subStr[i] != processStr[i] {
				return false
			}
		}
		return true
	}

	i := 0
	for i = pLen; i < sLen; i++ {
		if isAnagram(processStr, subStr) {
			ret = append(ret, i-pLen)
		}
		processStr[int(s[i])]++ //s[i] will be released after on cycle
		processStr[int(s[i-pLen])]--
	}

	if isAnagram(processStr, subStr) {
		ret = append(ret, i-pLen)
	}

	return ret
}

//timeout
// func findAnagrams(s string, p string) []int {
// 	pLen := len(p)
// 	sLen := len(s)
// 	ret := make([]int, 0, 10)
// 	if sLen == 0 || pLen > sLen {
// 		return ret
// 	}
// 	if pLen == 0 {
// 		for i := 0; i < sLen; i++ {
// 			ret = append(ret, i)
// 		}
// 		ret = append(ret, sLen)
// 		return ret
// 	}
//
// 	sumOfp, tmpOfp := 0, 0
// 	// CharListExist := make([]int, 256)
// 	// 	CharListExist[int(p[i])]++
// 	for i := 0; i < pLen; i++ {
// 		sumOfp += int(p[i])
// 		tmpOfp += int(s[i])
// 	}
// 	isAnagrams := func(s, p string, i int) bool {
// 		var innerMatch bool
// 		for j := 0; j < pLen; j++ {
// 			innerMatch = false
// 			for n := i; n < i+pLen; n++ {
// 				if p[j] == s[n] {
// 					innerMatch = true
// 				}
// 			}
// 			if !innerMatch {
// 				return false
// 			}
// 		}
// 		return true
// 	}
//
// 	for i := 0; i < sLen; i++ {
// 		if tmpOfp == sumOfp {
// 			if isAnagrams(s, p, i) {
// 				ret = append(ret, i)
// 			}
// 		}
// 		if i+pLen > sLen-1 {
// 			return ret
// 		}
// 		tmpOfp -= int(s[i])
// 		tmpOfp += int(s[i+pLen])
// 	}
// 	return ret
// }
