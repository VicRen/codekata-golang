package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Trim("Bob hit a ball, the hit BALL flew far after it was hit.", ",."))
}

func maxCountWord(s, ban string) string {
	s = strings.ToLower(s)
	sb := strings.Builder{}
	for _, cn := range s {
		c := byte(cn)
		if isAlphabet(c) || c == ' ' {
			sb.WriteByte(c)
		}
	}
	ws := strings.Split(sb.String(), " ")
	m := make(map[string]int)
	for _, w := range ws {
		if w == ban {
			continue
		}
		m[w]++
	}
	max := 0
	w := ""
	for k, v := range m {
		if v > max {
			max = v
			w = k
		}
	}
	return w
}

func maxCountWord2(s, ban string) string {
	m := make(map[string]int)
	s = strings.ToLower(s)
	sb := strings.Builder{}
	for _, cn := range s {
		c := byte(cn)
		if isAlphabet(c) {
			sb.WriteByte(c)
			continue
		}
		if c == ' ' {
			word := sb.String()
			if word != ban {
				m[word]++
			}
			sb = strings.Builder{}
			continue
		}
	}
	max := 0
	w := ""
	for k, v := range m {
		if v > max {
			max = v
			w = k
		}
	}
	return w

}

func isAlphabet(c byte) bool {
	return c >= 'a' && c <= 'z'
}
