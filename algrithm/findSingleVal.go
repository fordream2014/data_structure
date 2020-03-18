package algrithm

import "fmt"

func findSingleVal(arr []int) {
	if len(arr) == 0 {
		return
	}

	sum := 0
	for i:=0; i<len(arr); i++ {
		sum = sum ^ arr[i]
	}

	if sum == 0 {
		return
	}

	var num uint = 0
	for sum > 0 {
		sum = sum >> 1
		num++
		if sum & 1 == 1 {
			break
		}
	}

	fmt.Println(num)

	result0 := 0
	count0 := 0
	result1 := 0
	count1 := 0
	for i:=0; i<len(arr); i++ {
		temp := arr[i] >> num
		if temp & 1 == 1 {
			result1 ^= arr[i]
			count1 ++
		} else {
			result0 ^= arr[i]
			count0 ++
		}
	}
	
	fmt.Println(result0, count0, result1, count1)
	if count1 & 1 == 1 {
		fmt.Println(result1)
	} else {
		fmt.Println(result0)
	}

}

//TestFindSingleVal 数组中出现1次的数
func TestFindSingleVal() {

	//arr := []int{1,2,4,5,6,4,2}
	arr := []int{6,3,4,5,5,4,3}
	findSingleVal(arr)
}