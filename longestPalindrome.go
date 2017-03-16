package demo

//palindrom:
// b aba  ababababa
// ab abab
func longestPalindrome(s string) int {
	characterCount := make([]int, int('z')-int('A')+1)
	sLen := len(s)
	for i := 0; i < sLen; i++ {
		characterCount[int(s[i])-int('A')]++
	}
	characterCountLen := len(characterCount)
	count := 0
	single := false
	for i := 0; i < characterCountLen; i++ {
		count += characterCount[i]
		if characterCount[i]%2 == 1 {
			single = true
			count--
		}
	}
	if single {
		return count + 1
	}
	return count
}
