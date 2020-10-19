package main

import "fmt"

func main() {
	sum := 0
	for y := 1901; y <= 2000; y++ {
		for m := 1; m <= 12; m++ {
			if calWeek(y, m, 1) == 0 {
				sum++
			}
		}
	}
	fmt.Println("sum:", sum)
}

func calWeek(y, m, d int) int {
	//w=y+[y/4]+[c/4]-2c+[26(m+1)/10]+d-1
	//
	//公式中的符号含义如下，w：星期；c：世纪；y：年（两位数）；
	// m：月（m大于等于3，小于等于14，即在蔡勒公式中，某年的1、2月要看作上一年的13、14月来计算，
	// 比如2003年1月1日要看作2002年的13月1日来计算）；d：日；[ ]代表取整，即只要整数部分。
	if m == 1 || m == 2 {
		y--
		m += 12
	}
	c := y / 100
	y = y - c*100
	return (y + y/4 + c/4 - 2*c + 26*(m+1)/10 + d - 1) % 7
}

func isLeapYear(year int) bool {
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}
