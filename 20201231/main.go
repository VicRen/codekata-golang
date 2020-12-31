package main

import (
	"strconv"
	"strings"
)

func main() {

}

func process(n int) string {
	ret := ""
	for i := 1; i <= n; i++ {
		ret = say(ret)
	}
	return ret
}

func say(num string) string {
	if num == "" {
		return "1"
	}
	var nums []string
	var n []uint8
	last := num[0]
	for i := 0; i < len(num); i++ {
		if last == num[i] {
			n = append(n, num[i])
		} else {
			nums = append(nums, string(n))
			last = num[i]
			n = nil
			n = append(n, num[i])
		}
	}
	nums = append(nums, string(n))
	ret := strings.Builder{}
	for _, s := range nums {
		ret.WriteString(strconv.Itoa(len(s)))
		ret.WriteByte(s[0])
	}
	return ret.String()
}
