package main

func main() {

}

func isEcho(s string) bool {
	if s == "" {
		return false
	}
	l := 0
	r := len(s) - 1
	for r > l {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func isAlmostEcho(s string) bool {
	if s == "" {
		return false
	}
	l := 0
	r := len(s) - 1
	for r > l {
		if s[l] != s[r] {
			if !isEcho(removeIndex(l, s)) && !isEcho(removeIndex(r, s)) {
				return false
			}
			return true
		}
		l++
		r--
	}
	return true
}

func removeIndex(index int, s string) string {
	if index >= len(s) {
		return s
	}
	ret := []rune(s)
	ret = append(ret[0:index], ret[index+1:]...)
	return string(ret)
}
