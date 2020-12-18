package main

func main() {

}

func calculate(src []string) int {
	var cw string
	var max int
	for i := 0; i < len(src)-1; i++ {
		cw = src[i]
		cl := len(cw)
		for j := i + 1; j < len(src); j++ {
			cw2 := src[j]
			v := cl * len(cw2)
			if !hasSameChar(cw, cw2) && (v > max) {
				max = v
			}
		}
	}
	return max
}

func hasSameChar(s1, s2 string) bool {
	for _, v := range s1 {
		for _, v2 := range s2 {
			if v == v2 {
				return true
			}
		}
	}
	return false
}
