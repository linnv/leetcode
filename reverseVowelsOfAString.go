package demo

func reverseVowel(s string) string {
	var vowels = [...]bool{
		int('a'): true,
		int('e'): true,
		int('i'): true,
		int('o'): true,
		int('u'): true,
		int('A'): true,
		int('E'): true,
		int('I'): true,
		int('O'): true,
		int('U'): true,
	}
	bs := []byte(s)
	sLen := len(s)
	vowelIndex := make([]int, 0, sLen)
	for i := 0; i < sLen; i++ {
		if int(bs[i]) <= int('u') && vowels[int(bs[i])] {
			vowelIndex = append(vowelIndex, i)
		}
	}
	vowelIndexLen := len(vowelIndex)
	half := vowelIndexLen / 2
	for i := 0; i < half; i++ {
		bs[vowelIndex[i]], bs[vowelIndex[vowelIndexLen-i-1]] = bs[vowelIndex[vowelIndexLen-i-1]], bs[vowelIndex[i]]
	}
	return string(bs)
}
