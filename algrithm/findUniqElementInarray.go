package algrithm

import "fmt"

//在一个数组中除一个数字只出现一次之外，其他数字都出现了三次，请找出那个只出现一次的数字
func findUniqElementInarray(base []int) int {
	if len(base) == 0 {
		return 0
	}

	bitarr := make([]int, 32)
	for i:=0; i<len(base); i++ {
		bitmask := 1
		for j:=31; j>=0; j-- {
			bitval := base[i] & bitmask
			if bitval != 0 {
				bitarr[j] += 1
			}
			bitmask = bitmask << 1
		}
	}

	fmt.Println(bitarr)

	var result int = 0
	for i:=0; i<32; i++ {
		result = result << 1
		result += bitarr[i] % 3
	}

	return result
}

func TestFindUniqElementInarray() {
	arr := []int{2,2,2,4,3,3,3}
	res := findUniqElementInarray(arr)
	fmt.Println(res)
}