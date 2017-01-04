package demo

//like cycle linklist:`https://play.golang.org/p/Bj4al1ARXo`
func findDuplicate(nums []int) int {
	//why travel from the end?
	// because you can index on 0 but no n+1 for a array which length is n+1
	len := len(nums)
	speed1x := nums[len-1]
	speed2x := nums[nums[len-1]-1]
	for speed1x != speed2x {
		speed1x = nums[speed1x-1]
		speed2x = nums[nums[speed2x-1]-1]
	}
	//len in here just like head in: https://play.golang.org/p/Bj4al1ARXo
	speed2x = len
	for speed1x != speed2x {
		speed1x = nums[speed1x-1]
		speed2x = nums[speed2x-1]
	}
	return speed1x
}
