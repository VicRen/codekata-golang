package main

func main() {

}

func turnOver(s string) string {
	var letters []int32
	m := make(map[int]int32)
	for i, c := range s {
		if isLetter(c) {
			letters = append(letters, c)
		} else {
			m[i] = c
		}
	}
	var ret []int32
	for i := len(letters) - 1; i >= 0; i-- {
		ret = append(ret, letters[i])
	}
	for k, v := range m {
		ret = insert(ret, k, v)
	}
	return string(ret)
}

func isLetter(s int32) bool {
	return (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z')
}

func insert(a []int32, index int, value int32) []int32 {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}
