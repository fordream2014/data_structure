package algrithm

import "fmt"

//数组中有N+2个数，其中N个数出现了偶数次，2个数出现了奇数次，找出这两个数
func findOddTimesElement(arr []int) {
	//使用异或方法
	if len(arr) == 0 {
		return
	}

	result := 0
	for _,v := range arr {
		result = result ^ v
	}

	//result 为两个出现奇数次的数的异或结果
	position := uint(0)
	for i:=result; i&1==0; i=i>>1 {
		position ++
	}

	odd := 0
	for _,v := range arr {
		if (v >> position) & 1 == 0 {
			odd = odd ^ v
		}
	}

	fmt.Println("其中一个出现奇数次的数为", odd)
	fmt.Println("另一个出现奇数次的数为", result^odd)
}

func TestFindOddTimesElement() {
	arr := []int{3,5,6,6,5,7,2,2}
	findOddTimesElement(arr)
}
