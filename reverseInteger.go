package demo

import "strconv"

func reverse(x int) int {
	const (
		maxInt = 1<<31 - 1
		minInt = -maxInt - 1
	)
	y := 0
	for x != 0 {
		move := x % 10
		if y > maxInt/10 || y < minInt/10 {
			return 0
		}
		y = y*10 + move
		x /= 10
	}
	return y
}

//this costs more 10x time than upper func
func reverseByswapChart(x int) int {
	s := []byte(strconv.Itoa(x))
	sLen := len(s)
	if s[0] == '-' || s[0] == '+' {
		for i := 1; i <= sLen/2; i++ {
			s[sLen-i], s[i] = s[i], s[sLen-i]
		}
	} else {
		for i := 0; i < sLen/2; i++ {
			s[sLen-1-i], s[i] = s[i], s[sLen-1-i]
		}
	}
	// r, _ := strconv.Atoi(string(s))
	r := myAtoi(string(s))
	const (
		maxInt = 1<<31 - 1
		minInt = -maxInt - 1
	)
	if r < minInt || r > maxInt {
		return 0
	}
	return r
}
