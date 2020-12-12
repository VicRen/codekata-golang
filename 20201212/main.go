package main

func main() {

}

func compress(src []int) []int {
	var ret []int
	for i := 0; 2*i+1 < len(src); i++ {
		freq := src[2*i]
		val := src[2*i+1]
		for j := 0; j < freq; j++ {
			ret = append(ret, val)
		}
	}
	return ret
}
