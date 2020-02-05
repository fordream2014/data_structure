package algrithm

import "fmt"

func getLongestIncrSequenceSize(seq []int) int {
	if seq == nil || len(seq) == 0 {
		return 0
	}

	max := 1
	start := 0
	for i:=1; i<len(seq); i++ {
		if seq[i] > seq[i-1] {
			if i - start + 1 > max {
				max = i - start + 1
			}
		} else {
			start = i
		}
	}
	return max
}

func TestGetLongestIncrSequenceSize() {
	//nums := []int{1,3,5,4,7}
	nums := []int{2,2,2,2,2}
	max := getLongestIncrSequenceSize(nums)
	fmt.Println(max)
}
