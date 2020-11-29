package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	i := 10
	for {
		if isSameDigits([6]int{i, 2 * i, 3 * i, 4 * i, 5 * i, 6 * i}) {
			break
		}
		i++
	}
	fmt.Println("i=", i)
}

func isSameDigits(value [6]int) bool {
	last := 0
	for i, n := range value {
		x := sortInt(n)
		if i == 0 {
			last = x
		} else if x != last {
			return false
		}
	}
	return true
}

func sortInt(input int) int {
	swapped, _ := strconv.Atoi(bubbleSort(strconv.Itoa(input)))
	return swapped

}

func bubbleSort(word string) string {
	wordtable := strings.Split(word, "")
	for j := 0; j < len(word); j++ {

		for i := 0; i < len(word)-1; i++ {
			if wordtable[i] < wordtable[i+1] {
				temp := wordtable[i]
				wordtable[i] = wordtable[i+1]
				wordtable[i+1] = temp
			}
		}
	}
	return strings.Join(wordtable, "")
}
