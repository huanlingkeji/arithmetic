package locde

import (
	"fmt"
	"testing"
)

func removeOccurrences(s string, part string) string {
	lthp := len(part)
	for i := 0; i <= len(s)-lthp; {
		if s[i:i+lthp] == part {
			s = s[:i] + s[i+lthp:]
			i -= lthp
			if i < 0 {
				i = 0
			}
		} else {
			i++
		}
	}
	return s
}

func TestName(t *testing.T) {
	fmt.Println(removeOccurrences("daabcbaabcbc", "abc"))
}
