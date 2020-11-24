package main

import "strconv"

func main() {

}

func compress(str string) string {
	var cl int32
	count := 0
	var ret string
	for _, s := range str {
		if s != cl {
			cl = s
			count = 1
			ret += string(s) + "1"
		} else {
			count++
			ret = ret[:len(ret)-1] + strconv.Itoa(count)
		}
		if len(ret) >= len(str) {
			return str
		}
	}
	return ret
}
