package main

func main() {

}

func countBalance(s string) int {
	count, rCount, lCount := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'R' {
			rCount++
		} else if s[i] == 'L' {
			lCount++
		}
		if rCount > 0 && rCount == lCount {
			lCount, rCount = 0, 0
			count++
		}
	}
	return count
}
