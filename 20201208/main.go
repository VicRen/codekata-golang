package main

import "fmt"

func main() {
}

func canSeparate(src []int) bool {
	//A[0] + A[1] + ... + A[i] == A[i+1] + A[i+2] + ... + A[j-1] == A[j] + A[j-1] + ... + A[A.length - 1]
	l := len(src)
	for i := 1; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			s1 := sumSlice(src[:i])
			s2 := sumSlice(src[i:j])
			s3 := sumSlice(src[j:])
			if s1 == s2 && s2 == s3 {
				fmt.Println("separated on: ", i, j)
				return true
			}
		}
	}
	return false
}

func sumSlice(src []int) int {
	ret := 0
	for _, n := range src {
		ret += n
	}
	return ret
}
