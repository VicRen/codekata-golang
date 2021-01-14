package main

import "fmt"

func main() {

}

func readBinaryWatch(num int) []string {
	var ret []string
	for h := 0; h <= 12; h++ {
		hCount := 0
		s := fmt.Sprintf("%#b", h)
		for _, c := range s {
			if c == '1' {
				hCount++
			}
		}
		for m := 0; m <= 60; m++ {
			mCount := 0
			s := fmt.Sprintf("%#b", m)
			for _, c := range s {
				if c == '1' {
					mCount++
				}
			}
			if hCount+mCount == num {
				if m > 10 {
					ret = append(ret, fmt.Sprintf("%d:%d", h, m))
				} else {
					ret = append(ret, fmt.Sprintf("%d:0%d", h, m))
				}
			}
		}
	}
	return ret
}
