package main

import "testing"

func IntToString(n int) string {
	r := make([]byte, 0)
	m := false
	if n < 0 {
		m = true
		n = -n
	}
	for n > 0 {
		r = append(r, byte('0'+n%10))
		n /= 10
	}
	if m {
		r = append(r, '-')
	}
	for s, e := 0, len(r)-1; s < e; s, e = s+1, e-1 {
		r[s], r[e] = r[e], r[s]
	}
	return string(r)
}

func TestIntToString(t *testing.T) {
	tt := []struct {
		name  string
		input int
		want  string
	}{
		{"1", 1, "1"},
		{"12", 12, "12"},
		{"121", 121, "121"},
		{"121333", 121333, "121333"},
		{"-121333", -121333, "-121333"},
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
