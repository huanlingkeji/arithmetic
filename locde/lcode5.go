package locde

func longestPalindrome(s string) string {
	res := ""
	for i := range s {
		s1 := getPalindrome(s, i, i)
		s2 := getPalindrome(s, i, i+1)
		if len(s1) > len(res) {
			res = s1
		}
		if len(s2) > len(res) {
			res = s2
		}
	}
	return res
}

func getPalindrome(s string, l, r int) string {
	lth := len(s)
	for l >= 0 && r < lth && s[l] == s[r] {
		l--
		r++
	}
	if l+1 <= r {
		return s[l+1 : r]
	}
	return ""
}
