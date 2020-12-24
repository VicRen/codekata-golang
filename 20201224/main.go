package main

func main() {

}

func isEcho(s string) bool {
	var sc []uint8
	for i := 0; i < len(s); i++ {
		c := s[i]
		if isAlphabetAndDigit(c) {
			if c <= 'Z' && c >= 'A' {
				c += 'a' - 'A'
			}
			sc = append(sc, c)
		}
	}
	if len(sc) == 0 {
		return true
	}
	l, r := 0, len(sc)-1
	for {
		if l >= r {
			break
		}
		if sc[l] != sc[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func isAlphabetAndDigit(c uint8) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}
