package algrithm

import (
	"fmt"
	"math"
)

//找出数组中第k小的数 k从1开始计算
func findKMinElement(arr []int, start, end, k int) int {
	if len(arr) == 0 || end < start {
		return math.MinInt64
	}

	r := partition(arr, start, end)
	fmt.Println(r, arr)
	//关键是递归用的开始和结束节点位置
	if r-start+1 == k {
		return arr[r]
	} else if r-start+1 > k {
		return findKMinElement(arr, start, r-1, k)
	} else {
		return findKMinElement(arr, r+1, end, k-(r-start+1))
	}
}

//拆分
func partition(arr []int, start, end int) int {
	base := arr[end]
	index := start
	for i:=start; i<=end; i++ {
		if arr[i] < base {
			if i != index {
				arr[i], arr[index] = arr[index], arr[i]
			}
			index ++
		}
	}
	arr[index], arr[end] = arr[end], arr[index]
	return index
}

func TestFindKMinElement() {
	arr := []int{4,0,1,0,2,3}
	k := 4
	result := findKMinElement(arr, 0, len(arr)-1, k)
	fmt.Printf("数组中第%v小的数为%v", k, result)
}


