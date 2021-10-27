package search

//连续的空间线性搜索，这就是二分查找可以发挥作用的标志

/*
for (int i = 0; i < n; i++)
    if (isOK(i))
        return ans;
*/

func shipWithinDays(weights []int, days int) int {
	left, right := max(weights), sum(weights)+1
	for left < right {
		mid := left + (right-left)/2
		if canFinish(weights, days, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func canFinish(weights []int, days int, cap int) bool {
	i := 0
	for needDay := 0; needDay < days; needDay++ {
		maxCap := cap
		// 第i件物品能搬运  go的while写法的特殊性
		for maxCap -= weights[i]; maxCap >= 0; maxCap -= weights[i] {
			i++
			//完成
			if i == len(weights) {
				return true
			}
		}
	}
	return false
}

func max(weights []int) int {
	s := -1
	for _, v := range weights {
		if v > s {
			s = v
		}
	}
	return s
}

func sum(weights []int) int {
	s := 0
	for _, v := range weights {
		s += v
	}
	return s
}

func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		} else if target < arr[mid] {
			right = mid - 1
		} else if arr[mid] < target {
			left = mid + 1
		}
	}
	return -1
}

func LeftBound(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			right = mid - 1
		} else if target < arr[mid] {
			right = mid - 1
		} else if arr[mid] < target {
			left = mid + 1
		}
	}
	// 全部比目标小 or  目标不存在
	if left >= len(arr) || arr[left] != target {
		return -1
	}
	return left
}

func RightBound(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			left = mid + 1
		} else if target < arr[mid] {
			right = mid - 1
		} else if arr[mid] < target {
			left = mid + 1
		}
	}
	// 全部比目标大 or 该值不存在
	if right < 0 || arr[right] != target {
		return -1
	}
	return right
}
