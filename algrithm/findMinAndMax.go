package algrithm

import "fmt"

//查找数组中的最大值和最小值
//1 可以使用蛮力法，需要时间复杂度为O(n)，需要对比2n-2次
//2 使用分治法，时间复杂度为O(n)，需要对比次数：n/2 + 2(n/2) = 3n/2
func findMinAndMax(arr []int) (min, max int) {
	//使用分治法
	if len(arr) == 0 {
		return 0,0
	}
	if len(arr) == 1 {
		return arr[0], arr[0]
	}
	for i:=0; i<len(arr) - 1; i+=2 {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}
	fmt.Println(arr)

	//寻找最小值
	min = arr[0]
	for i:=2; i<len(arr)-1; i+=2 {
		if arr[i] < min {
			min = arr[i]
		}
	}

	max = arr[1]
	for i:=3; i<len(arr)-1; i+=2 {
		if arr[i] > max {
			max = arr[i]
		}
	}

	length := len(arr)
	if length % 2 == 1 {
		//奇数个
		last := arr[length - 1]
		if last > max {
			max = last
		}
		if last < min {
			min = last
		}
	}

	return min, max
}

func TestFindMinMax() {
	arr := []int{7,3,19,40,4,7,1}
	min,max := findMinAndMax(arr)
	fmt.Println(min, max)
}

