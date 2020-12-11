package main

import (
	"fmt"
	"strings"
)

//给定一个字符串 s 和一个整数 k，你需要对从字符串开头算起的每隔 2k 个字符的前 k 个字符进行反转。
//
//如果剩余字符少于 k 个，则将剩余字符全部反转。
//如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。

func main() {

}

func process(k int, s string) string {
	sb := strings.Builder{}
	l := 0
	r := 2 * k
	for {
		fmt.Println(l, r)
		if r > len(s) {
			if l > len(s) {
				sb.WriteString(s[r-2*k:])
			} else {
				sb.WriteString(reverse(s[l : l+k]))
				sb.WriteString(s[l+k:])
			}
			break
		}
		sb.WriteString(reverse(s[l : l+k]))
		sb.WriteString(s[l+k : r])
		l += 2 * k
		r += 2 * k
	}
	return sb.String()
}

func reverse(s string) string {
	src := []rune(s)
	l := len(src)
	for i := 0; i < l/2; i++ {
		src[i], src[l-i-1] = src[l-i-1], src[i]
	}
	return string(src)
}
