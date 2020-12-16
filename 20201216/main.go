package main

import "fmt"

func main() {

}

func process(n int) []string {
	m := make(map[string]struct{})
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			g := findGCD(i, j)
			m[fmt.Sprintf("%d/%d", j/g, i/g)] = struct{}{}
		}
	}
	var ret []string
	for k := range m {
		ret = append(ret, k)
	}
	return ret
}

func findGCD(a, b int) int {
	for a > 0 {
		t := a
		a = b % a
		b = t
	}
	return b
}
