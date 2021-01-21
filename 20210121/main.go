package main

import "strings"

const balloon = "balloon"

func main() {

}

func maxNumberOfBalloons(text string) int {
	nt := text
	lt := text
	end := false
	count := 0
	for {
		for _, c := range balloon {
			nt = strings.Replace(lt, string(c), "", 1)
			if nt == lt {
				end = true
				break
			}
			lt = nt
		}
		if end {
			break
		}
		count++
	}
	return count
}
