package main

import (
	"errors"
	"fmt"
)

func main() {
	nums, _ := findSprialDiagonals(1001)
	sum := 0
	for _, n := range nums {
		sum += n
	}
	fmt.Printf("nums=%v\nsum=%d\n", nums, sum)
}

func findSprialDiagonals(width int) ([]int, error) {
	if width%2 == 0 {
		return nil, errors.New("invalid width")
	}
	x := 3
	ret := []int{1}
	cw := 3
	d := 2
	for width/cw > 0 {
		ret = append(ret, []int{x, x + d, x + 2*d, x + 3*d}...)
		x = x + 3*d + cw + 1
		cw += 2
		d = cw - 1
	}
	return ret, nil
}
