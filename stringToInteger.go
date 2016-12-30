package demo

import "strings"

//step 1: remove space
//step 2: check prefix char is number and -,+ or not
//step 3: caculate the number and check the max result while caculating
//step 4: check the number is posive or negative, or calculate the last bit

func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	fixCharDot, fixCharSub, fixCharPlus := int('.'), int('-'), int('+')
	zero, night := int('0'), int('9')
	validPrefix := func(s string) string {
		for i := 0; i < len(s); i++ {
			_i := int(s[i])
			if _i == fixCharDot {
				return ""
			}
			if _i == fixCharSub || _i == fixCharPlus {
				continue
			}
			if zero <= _i && _i <= night {
				return s
			}
			return ""
		}
		return s
	}

	str = validPrefix(str)
	if len(str) < 1 {
		return 0
	}

	getPow := func(carry int) int64 {
		if carry < 1 {
			return 0
		}
		r := int64(1)
		for i := 0; i < carry; i++ {
			r = r * 10
		}
		return r
	}

	var ret int64
	carry := 0
	const maxBit = 31
	const maxInt = int(1<<maxBit - 1)
	var minIntCom int64 = 1 << maxBit
	firstChar := int(str[0])
	for i := len(str) - 1; i > 0; i-- {
		_i := int(str[i])
		if zero <= _i && _i <= night {
			if carry == 0 {
				ret = int64(str[i]) - int64(zero)
				carry++
				continue
			}
			ret = ret + (int64(str[i])-int64(zero))*getPow(carry)
			if ret >= int64(maxInt) {
				if firstChar == fixCharSub {
					if ret >= minIntCom {
						return 1<<maxBit - 2*(1<<maxBit)
					}
					return int(ret - 2*ret)
				}
				return int(maxInt)
			}
			carry++
			continue
		}
		ret = 0
		carry = 0
	}
	if firstChar == fixCharSub {
		return int(ret - 2*ret)
	}
	if firstChar == fixCharPlus {
		return int(ret)
	}

	if carry == 0 {
		return firstChar - zero
	}
	ret = ret + int64(firstChar-zero)*getPow(carry)
	if ret >= int64(maxInt) {
		return int(maxInt)
	}
	return int(ret)
}
