package main

import (
	"fmt"
	"strings"
)

func main() {

}

func isCompatible(a, b, c string) bool {
	if c == "" {
		return a == "" && b == ""
	}
	if len(c) != len(a)+len(b) {
		return false
	}
	var r []uint8
	l := -1
	var ra []uint8
	for i := 0; i < len(a); i++ {
		ca := a[i]
		for {
			l++
			cc := c[l]
			if cc == ca {
				ra = append(ra, ca)
				break
			}
			r = append(r, c[l])
		}
	}
	for i := l; i < len(c); i++ {
		r = append(r, c[i])
	}
	fmt.Println(string(ra))
	fmt.Println(string(r))
	return strings.Contains(string(r), b)
}
