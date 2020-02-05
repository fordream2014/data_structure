package algrithm

import (
	"fmt"
	"sort"
)

//这种方式不能很好的去重
//例如：
//数组：[-1 0 1 2 -1 -4]
//1 [-1 0]
//-1 [-1 2]
//-1 [0 1]
//时间复杂度为：O(n^2)
//空间复杂度为：O(n)
func threeSummaryByHash(nums []int) {
	if nums == nil || len(nums) < 3 {
		return
	}
	fmt.Println(nums)
	hash := make(map[int][]int)
	for i:=0; i<len(nums); i++ {
		for j:=i+1; j<len(nums); j++ {
			if hash[nums[j]] != nil {
				fmt.Println(nums[j], hash[nums[j]])
				hash[nums[j]] = nil
			} else {
				mask := 0 - nums[i] - nums[j]
				hash[mask] = append([]int{}, nums[i], nums[j])
			}
		}
	}
}

//时间复杂度：O(nlogn) + O(n^2)
func threeSummaryBySort(nums []int) {
	if nums == nil || len(nums) < 3 {
		return
	}
	sort.Sort(sort.IntSlice(nums))
	fmt.Println(nums)

	if nums[0] > 0 {
		return
	}

	for index:=0; index<len(nums); index++ {
		//这一步很重要，不然会有重复元素
		if index - 1 > 0 && nums[index-1] == nums[index] {
			continue
		}
		left := index + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[index] + nums[left] + nums[right]
			if sum == 0 {
				fmt.Printf("%v:%v, %v:%v, %v:%v \n",index, nums[index], left,nums[left], right,nums[right])
				for left + 1 < right && nums[left] == nums[left+1] {
					left ++
				}
				for right - 1 > left && nums[right-1] == nums[right] {
					right --
				}
				left ++
				right --
			} else if sum > 0 {
				right --
			} else {
				left ++
			}
		}
	}
}

func TestSummaryThreeNumber() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	threeSummaryByHash(nums)

	//threeSummaryBySort(nums)
}
