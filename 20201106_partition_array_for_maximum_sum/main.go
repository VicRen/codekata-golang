package main

import (
	"fmt"
	"math"
)

var arr = []int{1, 40, 24, 22, 12, 4, 5, 10, 62, 9, 55, 14, 39, 89}

func main() {
	fmt.Println("solution:", findMaxSum(4, arr))
}

func findMaxSum(k int, arr []int) int {
	l := len(arr)
	dp := make([]int, l+1)
	dp[0] = 0
	for i := 0; i <= l; i++ {
		max := -1

		for j := 1; j <= k && i-j >= 0; j++ {
			max = int(math.Max(float64(max), float64(arr[i-j])))
			dp[i] = int(math.Max(float64(dp[i]), float64(dp[i-j]+max*j)))
		}
	}
	fmt.Println(dp)
	return dp[l]
}
