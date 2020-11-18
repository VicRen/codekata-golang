package main

import "fmt"

func main() {
	fmt.Println(findMaxLine(65535))
}

func findMaxLine(n int32) int32 {
	var ret int32
	for {
		if findNumber(ret) > n {
			ret--
			return ret
		}
		ret++
	}
}

func findNumber(n int32) int32 {
	return n * (n + 1) / 2
}
