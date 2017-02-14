package demo

func groupAnagrams(strs []string) [][]string {
	strsLen := len(strs)
	groups := make([][]string, 0, strsLen/4)
	if strsLen < 1 {
		return groups
	}
	sortStr := make([]string, strsLen)
	sortString := func(str string) string {
		bs := []byte(str)
		//or you can implement interface of []byte to get a sorted string
		return string(QuickSort(bs))
	}
	for i := 0; i < strsLen; i++ {
		sortStr[i] = sortString(strs[i])
	}

	mapIndex := make(map[string]int, 10)
	lastIndex := 0
	for i := 0; i < strsLen; i++ {
		if _, ok := mapIndex[sortStr[i]]; !ok {
			groups = append(groups, []string{strs[i]})
			mapIndex[sortStr[i]] = lastIndex
			lastIndex++
			continue
		}
		groups[mapIndex[sortStr[i]]] = append(groups[mapIndex[sortStr[i]]], strs[i])
	}
	return groups
}

func QuickSort(bs []byte) []byte {
	quickSort(bs, 0, len(bs)-1)
	return bs
}

func partion(bs []byte, left, right int) int {
	pivotValue := bs[right]
	pivotValueInt := int(pivotValue)
	pivotIndex := left
	for i := left; i < right; i++ {
		if int(bs[i]) < pivotValueInt {
			bs[pivotIndex], bs[i] = bs[i], bs[pivotIndex]
			pivotIndex++
		}
	}
	bs[pivotIndex], bs[right] = bs[right], bs[pivotIndex]
	return pivotIndex
}

func quickSort(bs []byte, left, right int) {
	if left > right {
		return
	}
	pivot := partion(bs, left, right)
	quickSort(bs, left, pivot-1)
	quickSort(bs, pivot+1, right)
}

func GroupAnagrams(strs []string) [][]string {
	return groupAnagrams(strs)
}

func GroupAnagramsTimeout(strs []string) [][]string {
	return groupAnagramsTimeout(strs)
}

func groupAnagramsTimeout(strs []string) [][]string {
	strsLen := len(strs)
	groups := make([][]string, 0, strsLen/4)
	isAnagrams := func(s, p string) bool {
		chartCount := make([]int, 256)
		sLen := len(s)
		pLen := len(p)
		if pLen != sLen {
			return false
		}
		for i := 0; i < sLen; i++ {
			chartCount[int(s[i])]++
			chartCount[int(p[i])]--
		}
		for i := 0; i < sLen; i++ {
			if chartCount[int(s[i])] != 0 {
				return false
			}
		}
		return true
	}

	existAnagramAlready := func(p string) (groupeIndex int) {
		for i := 0; i < len(groups); i++ {
			if isAnagrams(groups[i][0], p) {
				return i
			}
		}
		return -1
	}

	for i := 0; i < strsLen; i++ {
		if j := existAnagramAlready(strs[i]); j >= 0 {
			groups[j] = append(groups[j], strs[i])
			continue
		}
		groups = append(groups, []string{strs[i]})
	}
	return groups
}
