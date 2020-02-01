package algrithm

import "fmt"

func findLongestNoDuplicatSubString(base string) string {
	if len(base) <= 1 {
		return base
	}

	letters := make([]int, 26)
	for i:=0; i<26; i++ {
		letters[i] = -1
	}

	letter := int(base[0] - 'a')
	letters[letter] = 0

	start_index := 0
	maxLen := 0
	preLen := 1
	for i:=1; i<len(base); i++ {
		char := int(base[i] - 'a')
		if letters[char] < 0 || i - letters[char] > preLen {
			preLen++
		} else {
			preLen = i - letters[char]
		}
		letters[char] = i
		//fmt.Println(preLen)
		if preLen > maxLen {
			maxLen = preLen
			start_index = i - preLen + 1
		}
	}

	return base[start_index:start_index + maxLen]
}

//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
func TestFindLongestNoDuplicatSubString() {
	base := "abcabcbb"
	str := findLongestNoDuplicatSubString(base)
	fmt.Println(base, str)
	base = "bbbbb"
	fmt.Println(base, findLongestNoDuplicatSubString(base))
	base = "pwwkew"
	fmt.Println(base, findLongestNoDuplicatSubString(base))
}
