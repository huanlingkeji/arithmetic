package substring

func checkInclusion(s1 string, s2 string) bool {
	cMap := make(map[byte]int)
	needMap := make(map[byte]int)
	for i := range s1 {
		needMap[s1[i]]++
	}
	left, right := 0, 0
	for right < len(s2) {
		cMap[s2[right]]++
		right++
		if len(s2[left:right]) >= len(s1) {
			if checkOK(cMap, needMap) {
				return true
			}
			cMap[s2[left]]--
			left++
		}
	}
	return false
}

//func checkOK(cMap map[byte]int, needMap map[byte]int) bool {
//	for k, v := range needMap {
//		if cMap[k] < v {
//			return false
//		}
//	}
//	return true
//}
