package demo

func reverseWords(s string) string {
	sLen := len(s)
	bs := []byte(s)
	i := 0
	for i < sLen {
		if bs[i] == ' ' {
			i++
			continue
		}
		j := i + 1
		for j < sLen && bs[j] != ' ' {
			j++
		}
		nextIndex := j //index of nearest space
		j--
		for i < j {
			bs[i], bs[j] = bs[j], bs[i]
			i++
			j--
		}
		i = nextIndex + 1 //next none space  char
	}
	return string(bs)
}
