package search

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
