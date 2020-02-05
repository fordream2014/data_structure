package algrithm

import "fmt"

func findKMaxElement(nums []int, k int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	if k < 0 || k >= len(nums) {
		return -1
	}
	left := 0
	right := len(nums) - 1
	for left <= right {
		r := partitionArray(nums, left, right)
		if r == k {
			return nums[r]
		}else if r > k {
			right = r - 1
		} else {
			left = r + 1
		}
	}
	return -1
}

func partitionArray(nums []int, left, right int) int {
	if left >= right {
		return left
	}
	base := nums[right]
	start := left
	for i:=left; i<=right; i++ {
		if nums[i] > base {
			if i != start {
				//交换值
				nums[i], nums[start] = nums[start], nums[i]
			}
			start ++
		}
	}
	nums[start], nums[right] = nums[right], nums[start]
	return start
}

func TestFindKMaxElement() {
	//nums := []int{3,2,1,5,6,4}
	k := 3
	nums := []int{3,2,3,1,2,4,5,5,6}
	num := findKMaxElement(nums, k)
	fmt.Println(num)
}