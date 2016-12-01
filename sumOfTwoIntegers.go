package demo

// `e.g. also work with negetive numbers which symbol is 1
//0110, 6
//0010, 2
// 	0100 ^, merged bits that has no carrying
// 	0100 &, 010 <<1,carrying one
// 		0000  ^, add bits that has no carrying
// 		1000  &,0100<<1,carrying one,check next carrying before merging
// 			1000 ^, merge bits that has no carrying
// 			0000 &,ok no more carrying to handle and the sum result is 1000
// `

func getSum(a, b int) int {
	carry := a & b
	bitNocarrying := a ^ b
	var tmp int
	for carry != 0 {
		carry = carry << 1
		tmp = carry
		carry = bitNocarrying & carry       //get next carrying
		bitNocarrying = bitNocarrying ^ tmp //we add the carring bit
	}
	return bitNocarrying
}
