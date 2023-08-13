package main

import "testing"

func IntToString(n int) string {
	ret := make([]byte, 0)
	for n > 0 {
		ret = append([]byte{byte('0' + n%10)}, ret...)
		n /= 10
	}
	return string(ret)
}

func TestIntToString(t *testing.T) {
	tt := []struct {
		name  string
		input int
		want  string
	}{
		{"1", 1, "1"},
		{"12345678", 12345678, "12345678"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := IntToString(tc.input)
			if got != tc.want {
				t.Errorf("IntToString(%d) = %s, want %s\n", tc.input, got, tc.want)
			}
		})
	}
}
