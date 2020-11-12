package main

func main() {

}

func sortFor(n int, src []int) []int {
	var find []int
	for _, x := range src {
		if x == n {
			find = append(find, x)
		} else {
			left = append(left, x)
		}
	}
	return append(find, left...)
}

func sortFor2(n int, src []int) []int {
	l := len(src)
	for i := l - 1; i >= 0; i-- {
		for j := l - 1; j >= 0; j-- {
			if src[i] != n && src[j] == n {
				tmp := src[i]
				src[i] = src[j]
				src[j] = tmp
			}
		}
	}
	return src
}
