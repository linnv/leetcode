package demo

//notice: right side symbol must be in reverse order to left side symbol
func isValid(s string) bool {
	sLen := len(s)
	if sLen%2 == 1 {
		return false
	}
	if sLen == 0 {
		return true
	}

	isRightSymbol := func(b byte) (byte, bool) {
		if b == ']' {
			return '[', true
		}
		if b == '}' {
			return '{', true
		}
		if b == ')' {
			return '(', true
		}
		return '0', false
	}

	if _, t := isRightSymbol(s[0]); t {
		return false
	}

	stackLeftSymbol := make([]byte, 0, sLen)
	stackLeftSymbol = append(stackLeftSymbol, s[0])
	for i := 1; i < sLen; i++ {
		if b, t := isRightSymbol(s[i]); t {
			// check symbol of reverse position is match
			if stackLeftSymbol[len(stackLeftSymbol)-1] != b {
				return false
			}
			//pop
			stackLeftSymbol = stackLeftSymbol[:len(stackLeftSymbol)-1]
		} else {
			//push
			stackLeftSymbol = append(stackLeftSymbol, s[i])
		}
	}
	if len(stackLeftSymbol) > 0 {
		return false
	}
	return true
}
