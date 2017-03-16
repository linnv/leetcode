package demo

func detectCapitalUser(word string) bool {
	upperMax, lowerMin := int('Z'), int('a')
	wordLen := len(word)
	if wordLen < 1 {
		return false
	}
	if wordLen == 1 {
		return true
	}
	allLowerLeft := false
	if int(word[0]) <= upperMax && int(word[1]) >= lowerMin {
		allLowerLeft = true
	}
	if int(word[0]) >= lowerMin {
		allLowerLeft = true
		if int(word[1]) <= upperMax {
			return false
		}
	}
	for i := 2; i < wordLen; i++ {
		if int(word[i]) >= lowerMin && !allLowerLeft {
			return false
		}
		if int(word[i]) <= upperMax && allLowerLeft {
			return false
		}
	}
	return true
}
