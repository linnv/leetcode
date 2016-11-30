package demo

import "strconv"

func fizzBuzz(n int) []string {
	str := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			str = append(str, "FizzBuzz")
			continue
		}
		if i%3 == 0 {
			str = append(str, "Fizz")
			continue
		}
		if i%5 == 0 {
			str = append(str, "Buzz")
			continue
		}
		str = append(str, strconv.Itoa(i))
	}

	return str
}
