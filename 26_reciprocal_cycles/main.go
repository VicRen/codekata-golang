package main

import "fmt"

func main() {
	max := 0
	d := 1
	for d < 1000 {
		n := findCycleLengthOfOneNth(d)
		if n > max {
			max = d
		}
		d++
	}
	fmt.Println("max:", max)
}

func findCycleLengthOfOneNth(n int) int {
	m := map[int]int{}
	a := 1
	t := 0
	for {
		if _, ok := m[a]; ok {
			break
		}
		m[a] = t
		a = a % n * 10
		if a == 0 {
			return t
		}
		t++
	}
	return t - m[a]
}
