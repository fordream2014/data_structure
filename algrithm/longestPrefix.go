package algrithm

import "fmt"

func findLongestPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	temp := strs[0]
	for i:=1; i<len(strs); i++ {
		if temp == "" || strs[i] == "" {
			break
		} else {
			length := 0
			if len(temp) > len(strs[i]) {
				length = len(strs[i])
			} else {
				length = len(temp)
			}
			var j = 0
			for ; j<length; j++ {
				if temp[j] != strs[i][j] {
					break
				}
			}
			temp = temp[0:j]
		}
	}
	return temp
}

//使用归并算法求解
func findLongestPrefixByMerge(strs []string) string {
	return mergeFind(strs, 0, len(strs)-1)
}

func mergeFind(strs []string, start, end int) string {
	//fmt.Println(start, end)
	if start == end {
		return strs[start]
	}
	r := start + (end-start)>>1
	left := mergeFind(strs, start, r)
	right := mergeFind(strs, r + 1, end)
	return merge(left, right)
}

func merge(left, right string) string {
	 if left == "" || right == "" {
	 	return ""
	 }

	 length := len(left)
	 if len(left) > len(right) {
	 	length = len(right)
	 }
	 str := ""
	 for i:=0; i<length; i++ {
		if left[i] == right[i] {
			str += string(left[i])
		} else {
			break
		}
	 }
	 return str
}


func TestFindLongestPrefix() {
	strs := []string{
		"abc",
		"abcfas",
		"abcsdfw",
		"abcvd",
	}
	//str := findLongestPrefix(strs)
	//fmt.Println(str)

	prefix := findLongestPrefixByMerge(strs)
	fmt.Println(prefix)
}
