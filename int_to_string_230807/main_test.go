package main

import "testing"

func IntToString(n int) string {
	ret := make([]rune, 0)
	for n > 0 {
		num := n % 10
		ret = append(ret, '0'+rune(num))
		n /= 10
	}
	return ReverseString(ret)
}

func ReverseString(input []rune) string {
	l := len(input)
	for i := 0; i < l/2; i++ {
		input[i], input[l-i-1] = input[l-i-1], input[i]
	}
	return string(input)
}

func TestIntToString(t *testing.T) {
	tt := []struct {
		name  string
		input int
		want  string
	}{
		{"1", 1, "1"},
		{"11", 11, "11"},
		{"112", 112, "112"},
		{"1123789", 1123789, "1123789"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := IntToString(tc.input)
			if got != tc.want {
				t.Errorf("IntToString(%d) = %s, want %s", tc.input, got, tc.want)
			}
		})
	}
}
