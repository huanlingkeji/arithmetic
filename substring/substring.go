package substring

//滑动窗口题目

/*
int left = 0, right = 0;

while (right < s.size()) {
    window.add(s[right]);
    right++;

    while (valid) {
        window.remove(s[left]);
        left++;
    }
}
*/

// match  满足的数量 当cmap[s[right]] >= needMap[s[right]]时 match++
// 当 left需要右移时 cmap[s[right]] < needMap[s[right]] 则match--
// 当 match == len(needMap) 时则满足条件

func minWindow(s string, t string) string {
	cmap := make(map[byte]int)
	left, right := 0, 0
	needMap := make(map[byte]int)
	for i := range t {
		needMap[t[i]]++
	}
	minStr := s + " "
	for right < len(s) {
		cmap[s[right]]++
		right++
		for checkOK(cmap, needMap) {
			if len(s[left:right]) < len(minStr) {
				minStr = s[left:right]
			}
			cmap[s[left]]--
			left++
		}
	}
	if len(minStr) > len(s) {
		return ""
	}
	return minStr
}

func checkOK(cmap map[byte]int, needMap map[byte]int) bool {
	for k, v := range needMap {
		if cmap[k] < v {
			return false
		}
	}
	return true
}
