package demo

// log.Printf("int('z'): %+v\n", int('z'))
// log.Printf("int('a'): %+v\n", int('a'))
func canConstruct(ransomNote string, magazine string) bool {
	m := make([]int, int('z')+1)
	for i := 0; i < len(magazine); i++ {
		m[int(magazine[i])]++
	}
	for i := 0; i < len(ransomNote); i++ {
		m[int(ransomNote[i])]--
		if m[int(ransomNote[i])] < 0 {
			return false
		}
	}

	for i := 0; i < (int('z') + 1); i++ {
		if m[i] < 0 {
			return false
		}
	}
	return true
}
