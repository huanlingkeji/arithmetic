package locde

func isValid(s string) bool {
	stack := ""
	for i := range s {
		if s[i] == '[' || s[i] == '(' || s[i] == '{' {
			stack += string(s[i])
		} else {
			if len(stack) > 0 && leftof(s[i]) == stack[len(stack)-1] {
				stack = stack[:len(stack)-2]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func leftof(c byte) byte {
	if c == ']' {
		return '['
	}
	if c == ')' {
		return '('
	}
	return '{'
}
