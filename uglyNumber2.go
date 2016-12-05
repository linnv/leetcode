package demo

func getMininalTwo(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func getMinimalThree(two, three, five int) int {
	if two > getMininalTwo(three, five) {
		return getMininalTwo(three, five)
	}
	return two
}

//Multiply a ugly number by a ugly number is aslo ugly,so we can get the minimal number from xNum*{2,3,5}(xNum is the moving index of ugly array,each basic{2,3,5} ugly number has its own moving index)
func nthUglyNumber(n int) int {
	uglySlice := make([]int, 1, 1+n)
	uglySlice[0] = 1
	var move2, move3, move5 int
	appendNumer := 1
	for i := 0; i < n-1; i++ {
		appendNumer = getMinimalThree(uglySlice[move2]*2, uglySlice[move3]*3, uglySlice[move5]*5)
		if appendNumer == uglySlice[move2]*2 {
			move2++
		}
		if appendNumer == uglySlice[move3]*3 {
			move3++
		}
		if appendNumer == uglySlice[move5]*5 {
			move5++
		}
		uglySlice = append(uglySlice, appendNumer)
	}

	return appendNumer
}
