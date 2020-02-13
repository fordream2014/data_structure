package algrithm

import "fmt"

var degree int
//求逆序度
func getDegreeOfInversion(arr []int, start, end int) {
	if start >= end {
		return
	}

	mid := start + (end - start) >> 1
	getDegreeOfInversion(arr, start, mid)
	getDegreeOfInversion(arr, mid + 1, end)

	right_size := end - mid
	if right_size <= 0{
		return
	} else {
		right_index := end
		left_index := mid

		var merge = make([]int, end-start+1)
		index := end-start
		for right_index > mid && left_index >= start {
			temp := 0
			if arr[right_index] > arr[left_index] {
				temp = arr[right_index]
				right_index--
				right_size--
			} else {
				degree += right_size
				temp = arr[left_index]
				left_index--
			}
			merge[index] = temp
			index--
		}

		if left_index >= start {
			for left_index >= start {
				merge[index] = arr[left_index]
				left_index --
				index--
			}
		}
		if right_index > mid {
			for right_index > mid {
				merge[index] = arr[right_index]
				right_index --
				index--
			}
		}

		for i:=0; i<len(merge); i++ {
			arr[start+i] = merge[i]
		}
		//fmt.Println(start, end, arr, merge)
	}
}

func TestGetDegreeOfInversion() {
	arr := []int{7,5,6,4,3}
	getDegreeOfInversion(arr, 0, len(arr)-1)

	fmt.Println(degree, arr)
}
