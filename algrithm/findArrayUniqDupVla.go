package algrithm

import "fmt"

//采用数据映射法
//根据题目的要求：数字1-1000放在含有1001个元素的数组中，其中只有唯一的一个元素值重复，其他数字均只出现一次。
//数组下标值范围为0-1000，值域范围1-1000，对于任意个数i，f(i)称为它的后继，i称为f(i)的前驱，0只有后继。
//对于重复的值，肯定有两个前驱点。
func findUniqDupVal(arr []int) int {
	length := len(arr)
	if length <= 1  {
		return -1
	}

	index := 0
	for arr[index] > 0 {
		if arr[index] >= length {
			return -1
		}

		arr[index] = arr[index]*(-1)
		index = arr[index]*(-1)

		if index > length - 1 {
			return -1
		}
	}

	return index
}

//对于给定的自然数N，有一个N+M个元素的数组，其中存放了小于等于N的所有自然数，求重复出现的自然数序列{X}
func findMoreDupVals(arr []int, limit int) []int {
	if len(arr) <= 1  {
		return nil
	}
	length := len(arr)

	index := 0
	var result []int
	num := length - limit + 2
	for {
		if arr[index] < 0 {
			//重复值
			result = append(result, index)
			arr[index] = num
			num ++
		}
		if num == length {
			return result
		}

		arr[index] = arr[index]*(-1)
		index = arr[index]*(-1)

		if index > length - 1 {
			return nil
		}
	}
}

func TestFindUniqDupVal() {
	arr := []int{1,3,4,2,5,2}
	index := findUniqDupVal(arr)
	fmt.Println("唯一的重复值为: ", index)
}

func TestFindMoreDupVals() {
	//arr := []int{1,2,3,3,3,4,5,5,5,5,6}
	arr := []int{1,2,3,3,4,5,6,6,6,6,6}
	dups := findMoreDupVals(arr, 6)
	fmt.Println(dups)
}
