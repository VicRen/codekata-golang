package main

var r2d = map[uint8]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func main() {

}

func romaToDigit(s string) int {
	l := len(s)
	ret := 0
	for i := 0; i < l; i++ {
		c := r2d[s[i]]
		ret += c
		if i >= l-1 {
			break
		}
		cn := r2d[s[i+1]]
		if cn > c {
			ret -= 2 * c
		}
	}
	return ret
}
