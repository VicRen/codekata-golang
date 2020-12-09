package main

func main() {

}

func transform(src [][]int) [][]int {
	for i := 0; i < len(src); i++ {
		for j := i; j < len(src); j++ {
			src[i][j], src[j][i] = src[j][i], src[i][j]
		}
	}
	for i, l := range src {
		src[i] = reverse(l)
	}
	return src
}

func reverse(src []int) []int {
	l := len(src)
	for i := 0; i < l/2; i++ {
		src[i], src[l-i-1] = src[l-i-1], src[i]
	}
	return src
}
