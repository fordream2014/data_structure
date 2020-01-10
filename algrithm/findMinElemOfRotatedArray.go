package algrithm

import "fmt"

//找到旋转数组的最小元素
func findMinElemOfRotatedOfArray(arr []int, low, high int) int {
	if low == high {
		return arr[low]
	}

	mid := low + (high - low)/2
	if mid+1 <= high && arr[mid] > arr[mid+1] {
		return arr[mid+1]
	} else if mid-1>= low && arr[mid] < arr[mid-1]{
		return arr[mid]
	} else if arr[mid] > arr[high] {
		return findMinElemOfRotatedOfArray(arr, mid+1, high)
	} else if arr[mid] < arr[low] {
		return findMinElemOfRotatedOfArray(arr, low, mid-1)
	} else {
		// arr[mid] == arr[high] && arr[mid] == arr[low] 极端情况[1 1 0 1]
		leftMin := findMinElemOfRotatedOfArray(arr, low, mid-1)
		rightMin := findMinElemOfRotatedOfArray(arr, mid+1, high)
		if leftMin < rightMin {
			return leftMin
		} else {
			return rightMin
		}
	}
}

func TestFindMinElemOfRotatedOfArray() {
	arr := []int{5,6,7, 1,2,3,4}
	result := findMinElemOfRotatedOfArray(arr, 0, len(arr)-1)
	fmt.Println(result)

	arr = []int{1,1,0,1}
	result = findMinElemOfRotatedOfArray(arr, 0, len(arr)-1)
	fmt.Println(result)
}