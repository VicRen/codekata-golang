package main

func main() {

}

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	var stack []uint8
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch c {
		case '(':
			fallthrough
		case '{':
			fallthrough
		case '[':
			stack = append(stack, c)
		default:
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			if (last == '(' && c == ')') || last == '[' && c == ']' || last == '{' && c == '}' {
				stack = stack[:len(stack)-1]
			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
