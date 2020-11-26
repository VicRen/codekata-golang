package main

func main() {

}

func findMaxSubLen(src string) int {
	l := 0
	r := 0
	max := 0
	sm := make(map[uint8]struct{})
	for {
		if l == len(src) {
			return max
		}
		for {
			if r == len(src) {
				return max
			}
			c := src[r]
			if _, found := sm[c]; found {
				sm = make(map[uint8]struct{})
				l++
				r = l
				break
			} else {
				sm[c] = struct{}{}
				if len(sm) > max {
					max = len(sm)
				}
				r++
			}
		}
	}
}

func findMaxSubLen2(src string) int {
	l := 0
	r := 0
	max := 0
	var ss []uint8
	for {
		if l == len(src) {
			return max
		}
		for {
			if r == len(src) {
				return max
			}
			c := src[r]
			if i := indexOfValue(c, ss); i != -1 {
				ss = ss[i:]
				l += i + 1
				r++
				break
			} else {
				ss = append(ss, c)
				if len(ss) > max {
					max = len(ss)
				}
				r++
			}
		}
	}
}

func indexOfValue(n uint8, src []uint8) int {
	for i := len(src) - 1; i >= 0; i-- {
		if src[i] == n {
			return i
		}
	}
	return -1
}
