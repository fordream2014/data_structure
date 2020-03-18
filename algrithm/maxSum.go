package algrithm

import "fmt"

func maxSum(arr []int) {
	if len(arr) == 0 {
		return
	}

	max := 0
	sum := 0
	start := 0
	var result []int
	for i:=0; i<len(arr); i++ {
		if sum < 0 {
			start = i
			sum = 0
		}
		sum += arr[i]

		if sum > max {
			max = sum
			result = arr[start:i+1]
		}
	}

	fmt.Println(result)

}

func TestMaxSum() {
	arr := []int{1,-2,4,8,-4,7,-1,-5}
	maxSum(arr)
}