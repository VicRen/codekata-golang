package main

import (
	"strings"
)

func main() {

}

func transform(s string) []string {
	return process(strings.Split(s, " "))
}

func process(src []string) []string {
	maxL := 0
	for _, s := range src {
		if len(s) > maxL {
			maxL = len(s)
		}
	}
	mid := make([]string, maxL)
	for i := 0; i < len(mid); i++ {
		for _, s := range src {
			if i < len(s) {
				mid[i] += string(s[i])
			} else {
				mid[i] += " "
			}
		}
	}
	var ret []string
	for _, s := range mid {
		ret = append(ret, strings.TrimRight(s, " "))
	}
	return ret
}
