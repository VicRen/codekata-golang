package main

import "strings"

func main() {

}

func solution(s string) string {
	var stack [26]int
	max := 0
	for i := 0; i < len(s); i++ {
		index := s[i] - 'a'
		stack[index]++
		if stack[index] > max {
			max = stack[index]
		}
	}
	sb := strings.Builder{}
	up := true
	for i := 0; i < max; i++ {
		if up {
			for j := 0; j < len(stack); j++ {
				if stack[j] > 0 {
					sb.WriteByte(uint8(j) + 'a')
					stack[j]--
				}
			}
		} else {
			for j := len(stack) - 1; j >= 0; j-- {
				if stack[j] > 0 {
					sb.WriteByte(uint8(j) + 'a')
					stack[j]--
				}
			}
		}
		up = !up
	}
	return sb.String()
}
