package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

const BASE = 64

func main() {
	b, err := ioutil.ReadFile("names.txt")
	if err != nil {
		panic(err)
	}

	names := readNames(string(b))
	sort.Strings(names)
	fmt.Println(names)
	sum := 0
	for n, name := range names {
		sum += (n + 1) * worthOf(name)
	}
	fmt.Println("sum:", sum)
}

func readNames(s string) []string {
	if len(s) == 0 {
		return nil
	}
	str := strings.Replace(s, "\"", "", -1)
	ret := strings.Split(str, ",")
	return ret
}

func worthOf(word string) int {
	if len(word) == 0 {
		return 0
	}
	sum := 0
	for i := 0; i < len(word); i++ {
		a := word[i : i+1][0]
		w := uint8(a) - BASE
		sum += int(w)
	}
	return sum
}
