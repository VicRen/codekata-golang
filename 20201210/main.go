package main

func main() {

}

func replace(src string) string {
	var ret []rune
	l := len(src)
	cursor := 0
	for {
		if cursor == l {
			break
		}
		if cursor < l-1 && src[cursor] == '(' && src[cursor+1] == ')' {
			ret = append(ret, 'o')
			cursor += 2
			continue
		}
		if cursor < l-3 && src[cursor] == '(' && src[cursor+3] == ')' {
			ret = append(ret, []rune{'a', 'l'}...)
			cursor += 4
			continue
		}
		ret = append(ret, rune(src[cursor]))
		cursor++
	}
	return string(ret)
}
