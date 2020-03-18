package algrithm

import "fmt"

//找出数组中重复的数字
// 在一个长度为n的数组中，所有的数字都在0～n-1的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道
//每个数字重复了几次。请找出数组中任意一个重复的数字
func findArrayDupVals(digs []int) int {
	 if len(digs) == 0 {
	 	return -1
	 }

	 i := 0
	 for i < len(digs) {
	 	if i == digs[i] {
	 		i++
	 		continue
		} else {
			j := digs[i]
			if digs[j] == j {
				return j
			} else {
				digs[j] = digs[i]
				digs[i] = j
			}
		}
	 }

	 return -1
}

func TestFindArrayDupVals() {
	arr := []int{0,0,1,0,2,5,3}
	val := findArrayDupVals(arr)
	fmt.Println("结果", val)
}
