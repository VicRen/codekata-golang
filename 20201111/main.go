package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func main() {
	r := rand.Intn(100)
	fmt.Println("r:", r)
	var src []int
	for i := 0; i < 100; i++ {
		src = append(src, i)
		if i == r {
			continue
		}
		src = append(src, i)
	}
	n, _ := findSingleNumber(src)
	fmt.Println("lost number:", n)
	n, _ = findSingleNumber2(src)
	fmt.Println("lost number2:", n)
	n, _ = findSingleNumber3(src)
	fmt.Println("lost number3:", n)
}

func findSingleNumber(src []int) (int, error) {
	m := map[int]int{}
	for _, n := range src {
		m[n] += 1
	}
	for k, v := range m {
		if v == 1 {
			return k, nil
		}
	}
	return 0, errors.New("not found")
}

func findSingleNumber2(src []int) (int, error) {
	m := map[int]struct{}{}
	for _, n := range src {
		if _, ok := m[n]; ok {
			delete(m, n)
		} else {
			m[n] = struct{}{}
		}
	}
	for k := range m {
		return k, nil
	}
	return 0, errors.New("not found")
}

func findSingleNumber3(src []int) (int, error) {
	num := 0
	for _, n := range src {
		num = num ^ n
	}
	if num == 0 {
		return 0, errors.New("not found")
	}
	return num, nil
}
