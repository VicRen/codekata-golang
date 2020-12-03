package main

func main() {

}

func goodPairsCount(src []int) int {
	count := 0
	l := len(src)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if src[i] == src[j] {
				count++
			}
		}
	}
	return count
}

func goodPairsCount2(src []int) int {
	m := make(map[int]int)
	for _, n := range src {
		m[n] += 1
	}
	count := 0
	for _, v := range m {
		count += v * (v - 1) / 2
	}
	return count
}
