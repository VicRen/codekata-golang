package main

import "fmt"

var prices = []int{6, 8, 12, 14, 18, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44}
var pricesWithAll = []int{6, 8, 12, 12, 12, 12, 12, 14, 14, 18, 22, 22, 22, 24, 24, 26, 26, 26, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44}

func main() {
	fmt.Println(sumCoins(50, prices))
	fmt.Println(sumCoins(100, prices))
	fmt.Println("paid = 50", sumCoins(50, pricesWithAll))
	fmt.Println("paid = 100", sumCoins(100, pricesWithAll))
}

func sumCoins(target int, prices []int) int {
	t := make([]int, target+1)
	t[0] = 1
	for _, c := range prices {
		for j := c; j <= target; j++ {
			t[j] += t[j-c]
		}
	}
	return t[target]
}
