package main

import "fmt"

func main() {
	// d234 is divisible by 2 -> d4%2 == 0
	// d345 is divisible by 3 -> (d3+d4+d5)%3 == 0
	// d456 is divisible by 5 -> d6 == 5 || d6 == 0
	// d567 is divisible by 7
	// d678 is divisible by 11 -> d6 == 5 (if d6 == 0, d8 == d9) -> d678 = []int{506,517,528,539,561,572,583,594}
	// d789 is divisible by 13 -> d6789 = []int{5286,5390,5728,5832}
	// d8910 is divisible by 17 -> d678910 = []int{52867,53901,57289}
	// d5678910 = []int{952867,357289}
	// d345678910 = []int{30952867, 60357289,06357289}
	// left digits = []int{1,4}
	fmt.Println("sum:", 1430952867+1460357289+1406357289+4130952867+4160357289+4106357289)
}
