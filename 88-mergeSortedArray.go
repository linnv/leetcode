package demo

func MergeSortedArray(num1, num2 []int) []int {
	num1Len := len(num1)
	num2Len := len(num2)
	ret := make([]int, num2Len+num1Len)
	index1, index2 := num1Len-1, num2Len-1
	indexRet := num2Len + num1Len - 1
	for index1 > -1 && index2 > -1 {
		if num1[index1] == num2[index2] {
			ret[indexRet] = num1[index1]
			index1--
			index2--
			indexRet--
			continue
		}
		if num1[index1] > num2[index2] {
			ret[indexRet] = num1[index1]
			index1--
		} else {
			ret[indexRet] = num2[index2]
			index2--
		}
		indexRet--
	}

	for index1 > -1 {
		ret[indexRet] = num1[index1]
		index1--
		indexRet--
	}
	for index2 > -1 {
		ret[indexRet] = num2[index2]
		index2--
		indexRet--
	}
	if indexRet < 0 {
		return ret
	}
	return ret[indexRet+1:]
}
