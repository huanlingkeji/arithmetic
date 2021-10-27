package locde

func findDisappearedNumbers(nums []int) []int {
	byteMap := make(map[int]struct{})
	res := make([]int, 0)
	for _, v := range nums {
		byteMap[v] = struct{}{}
	}
	lth := len(nums)
	for i := 0; i < lth; i++ {
		if _, ok := byteMap[i]; !ok {
			res = append(res, i)
		}
	}
	return res
}
