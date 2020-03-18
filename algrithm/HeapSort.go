package algrithm

import "fmt"

//堆排序
//堆，完全二叉树，使用数组存储，元素从1开始
func heapSort(arr []int) {
	max := len(arr) - 1
	begin := max/2
	for begin > 0 {
		downSort(arr, begin)
		begin--
	}
}

func inserNode(arr []int, val int) []int {
	arr = append(arr, val)
	maxInd := len(arr) - 1

	parentInd := maxInd/2
	for parentInd > 0 {
		if arr[maxInd] < arr[parentInd] {
			arr[maxInd], arr[parentInd] = arr[parentInd], arr[maxInd]
			maxInd = parentInd
			parentInd = parentInd/2
		} else {
			break
		}
	}
	return arr
}

func deleteHeap(arr []int) int {
	max := len(arr)-1

	result := arr[1]
	arr[1], arr[max] = arr[max], arr[1]
	max--

	begin := 1
	for begin <= max/2 {
		left := begin * 2
		right := begin*2 + 1
		min := arr[left]
		if right <= max && arr[left] > arr[right] {
			min = arr[right]
		}
		if min > arr[begin] {
			break
		} else {
			if arr[left] == min {
				arr[begin], arr[left] = arr[left], arr[begin]
				begin = left
			} else {
				arr[begin], arr[right] = arr[right], arr[begin]
				begin = right
			}
		}
	}

	return result
}

func downSort(arr []int, begin int) {
	max := len(arr) - 1
	end := max/2
	for begin <= end {
		left_i := 2*begin
		right_i := 2*begin + 1
		//子节点最小值
		min_child := arr[left_i]
		if arr[left_i] > arr[right_i] {
			min_child = arr[right_i]
		}

		if arr[begin] <= min_child {
			break
		} else {
			if arr[left_i] == min_child {
				arr[begin], arr[left_i] = arr[left_i], arr[begin]
				begin = left_i
			} else {
				arr[begin], arr[right_i] = arr[right_i], arr[begin]
				begin = right_i
			}
		}
	}
}

//测试堆排序
func TestHeapSort() {
	arr := []int{0,4,6,8,5,9}
	heapSort(arr)

	fmt.Println(arr)
}