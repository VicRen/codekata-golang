package main

import (
	"math/big"
	"strings"
)

func main() {

}

func isLychrelNumber(n int) bool {
	return false
}

func IsPalindrome(n *big.Int) bool {
	str := n.String()
	pb := strings.Builder{}
	for i := len(str); i > 0; i-- {
		pb.WriteString(str[i-1 : i])
	}
	pstr := pb.String()
	return str == pstr
}
