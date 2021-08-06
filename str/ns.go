package str

import "strings"

func ns(s string) string {
	lth := len(s)
	i := lth - 1
	t := false
	for ; i > 1; i-- {
		c := s[i]
		if !t && c >= 'A' && c <= 'Z' {
			t = true
		}
		if t && 'a' <= c && c <= 'z' {
			s = s[:i+1] + "_" + s[i+1:]
			t = false
		}
	}
	return strings.ToLower(s)
}
