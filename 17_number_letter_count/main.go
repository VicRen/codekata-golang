package main

import "fmt"

func main() {
	count := 0
	for i := 1; i <= 1000; i++ {
		count += CountNumberLetter(i)
	}
	fmt.Println("count:", count)
}

var numberToLetter = map[int]string{
	1:    "one",
	2:    "two",
	3:    "three",
	4:    "four",
	5:    "five",
	6:    "six",
	7:    "seven",
	8:    "eight",
	9:    "nine",
	10:   "ten",
	11:   "eleven",
	12:   "twelve",
	13:   "thirteen",
	14:   "fourteen",
	15:   "fifteen",
	16:   "sixteen",
	17:   "seventeen",
	18:   "eighteen",
	19:   "nineteen",
	20:   "twenty",
	30:   "thirty",
	40:   "forty",
	50:   "fifty",
	60:   "sixty",
	70:   "seventy",
	80:   "eighty",
	90:   "ninety",
	100:  "hundred",
	1000: "thousand",
}

func CountNumberLetter(input int) int {
	n := input
	num := ""
	if input <= 20 {
		num += numberToLetter[input]
		fmt.Println("input:", n, "num:", num, len(num))
		return len(num)
	}
	if input/1000 > 0 {
		num += "onethousand"
		input = input % 1000
	}
	hundred := input / 100
	if hundred != 0 {
		num += numberToLetter[hundred] + "hundred"
	}
	input = input % 100
	one := input % 10
	ten := input / 10
	if one == 0 && hundred == 0 {
		num += numberToLetter[input]
	} else {
		if hundred > 0 && (one > 0 || ten > 0) {
			num += "and"
		}
		if ten == 1 {
			num += numberToLetter[input]
		} else {
			num += numberToLetter[ten*10] + numberToLetter[one]
		}
	}
	fmt.Println("input:", n, "num:", num, "len:", len(num))
	return len(num)
}
