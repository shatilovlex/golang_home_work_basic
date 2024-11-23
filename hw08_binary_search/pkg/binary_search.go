package pkg

func BinarySearch(arr []int, item int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		middle := left + (right-left)/2
		if arr[middle] >= item {
			right = middle - 1
		}
		if arr[middle] <= item {
			left = middle + 1
		}
		if arr[middle] == item {
			return middle
		}
	}
	return -1
}
