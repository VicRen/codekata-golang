package main

import "testing"

func IntToString(n int) string {
	var ret []byte
	for n > 0 {
		ret = append(ret, '0'+byte(n%10))
		n /= 10
	}
	return ReverseString(string(ret))
}

func IntToString2(n int) string {
	var ret []byte
	for n > 0 {
		ret = append([]byte{'0' + byte(n%10)}, ret...)
		n /= 10
	}
	return string(ret)
}

func ReverseString(s string) string {
	ret := []rune(s)
	for l, r := 0, len(ret)-1; l < r; l, r = l+1, r-1 {
		ret[l], ret[r] = ret[r], ret[l]
	}
	return string(ret)
}

func TestReverseString(t *testing.T) {
	tt := []struct {
		name  string
		input string
		want  string
	}{
		{"test", "test", "tset"},
		{"testing", "testing", "gnitset"},
		{"i am testing", "i am testing", "gnitset ma i"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := ReverseString(tc.input)
			if got != tc.want {
				t.Errorf("ReverseString(%s) = %s, want %s\n", tc.input, got, tc.want)
			}
		})
	}
}

func TestIntToString(t *testing.T) {
	tt := []struct {
		name  string
		input int
		want  string
	}{
		{"1", 1, "1"},
		{"12", 12, "12"},
		{"123", 123, "123"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			//got := IntToString(tc.input)
			got := IntToString2(tc.input)
			if got != tc.want {
				t.Errorf("IntToString(%d) = %s, want %s\n", tc.input, got, tc.want)
			}
		})
	}
}
