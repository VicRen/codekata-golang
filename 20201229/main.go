package main

func main() {

}

func isEqual(s, t string) bool {
	return process(s) == process(t)
}

func process(s string) string {
	var ret []rune
	for _, c := range s {
		if string(c) == "#" {
			if len(ret) > 0 {
				ret = ret[:len(ret)-1]
			}
			continue
		}
		ret = append(ret, c)
	}
	return string(ret)
}

func process2(s string) string {
	return ""
}
