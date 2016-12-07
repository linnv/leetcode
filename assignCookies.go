package demo

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	var j, count int
	for i := 0; i < len(g); i++ {
		for j < len(s) {
			if s[j] >= g[i] {
				count++
				j++
				break
			}
			j++
		}
	}
	return count
}
