package algrithm

import (
		"math"
	"fmt"
)

//找出最长回文子串
//暴力法 时间复杂度为O(n^3)
func findLongestSubStringByForce(base string) string {
	if len(base) < 2 {
		return base
	}

	maxLen := len(base)
	//定义子串开始和结束索引
	start := 0
	end := 0
	for i:=0; i<maxLen; i++ {
		for j:=i; j<maxLen; j++ {
			if isHuiwen(base, i, j) && (j-i)>(end-start){
				start = i
				end = j
			}
			//fmt.Printf("%v %v %v %v %v \n", base[i:j+1], i, j, start, end)
		}
	}
	//fmt.Println(start, end)
	return base[start:end+1]
}

//判断子串是否是回文
func isHuiwen(base string, start, end int) bool {
	if start > end {
		return false
	}
	if start == end {
		return true
	}
	//temp_start := start
	//temp_end := end
	for start <= end {
		if base[start] != base[end] {
			//fmt.Printf("%v验证结果：%v\n", base[temp_start:temp_end+1], false)
			return false
		}
		start++
		end--
	}
	//fmt.Printf("%v验证结果：%v\n", base[temp_start:temp_end+1], true)
	return true
}

//动态规划
//时间复杂度为O(n^2), 空间复杂度为O(n^2)
func findLongestSubStringByDP(base string) string {
	maxLen := len(base)
	//初始化
	dist := make([][]int, maxLen)
	for i:=0; i<maxLen; i++ {
		dist[i] = make([]int, maxLen)
	}

	start := 0
	end := 0
	for level:=0; level<maxLen; level++ {
		i := 0
		j := i + level
		for i<maxLen && j<maxLen{
			temp_dist := 0
			if level == 0 {
				temp_dist = 1
			} else if level == 1 && base[i] == base[j] {
				temp_dist = 1
			} else {
				if dist[i+1][j-1] == 1 && base[i] == base[j] {
					temp_dist = 1
				}
			}
			if temp_dist == 1 && (j-i)>(end-start){
				start = i
				end = j
			}
			dist[i][j] = temp_dist
			i++
			j = i+level
		}
	}

	return base[start:end+1]
}

//解决动态规划的空间复杂度问题
//采用中心扩展算法
func findLongestSubStringByMidExtend(base string) string {
	maxLen := len(base)
	left := 0
	right := 0
	for i:=0;i<maxLen;i++ {
		oddLen := extentFromMid(base, i, i)
		evenLen := extentFromMid(base, i, i+1)
		max := int(math.Max(float64(oddLen), float64(evenLen)))
		if max > right - left - 1 {
			if oddLen > evenLen {
				left = i - oddLen/2
				right = i + oddLen/2
			} else {
				left = i - evenLen/2 + 1
				right = i + evenLen/2
			}
		}
	}
	return base[left:right+1]
}

func extentFromMid(base string, left, right int) int {
	maxLen := len(base)
	for left >= 0 && right < maxLen && base[left] == base[right] {
		left--
		right++
	}
	return right-left-1
}

/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"
 */
func TestFindLongestHuiwen() {

	//暴力法
	//str := "babad"
	//substr := findLongestSubStringByForce(str)
	//fmt.Println(substr)
	//
	//str = "cbbd"
	//substr = findLongestSubStringByForce(str)
	//fmt.Println(substr)

	//动态规划
	str := "cbbd"
	//substr := findLongestSubStringByDP(str)
	//fmt.Println(substr)
	substr := findLongestSubStringByMidExtend(str)
	fmt.Println(substr)
}
