package algrithm

import "fmt"

//判断base中是否包含str的一个组合
func validateContainRank(base, str string) bool {
	str_len := len(str)
	base_len := len(base)
	if str_len == 0 || base_len == 0 {
		return false
	}
	if base_len < str_len {
		return false
	}

	str_chars := []byte(str)
	str_map := make([]int, 26)
	base_map := make([]int, 26)
	for _,v := range str_chars {
		ind := int(v - 'a')
		str_map[ind]++
	}

	for i:=0; i<base_len; i++ {
		ind := int(base[i] - 'a')
		base_map[ind]++
		if i >= str_len {
			char := int(base[i-str_len] - 'a')
			base_map[char]--
		}
		if validateMap(base_map, str_map) {
			return true
		}
	}
	return false
}

func validateMap(base, str []int) bool{
	for i:=0; i<26; i++ {
		if base[i] != str[i] {
			return false
		}
	}
	return true
}

/**
给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。
换句话说，第一个字符串的排列之一是第二个字符串的子串。

示例1:
输入: s1 = "ab" s2 = "eidbaooo"
输出: True
解释: s2 包含 s1 的排列之一 ("ba").

示例2:
输入: s1= "ab" s2 = "eidboaoo"
输出: False
注意：
输入的字符串只包含小写字母
两个字符串的长度都在 [1, 10,000] 之间
 */
func TestValidateContainRank() {
	s1 := "ab"
	s2 := "eidbaooo"
	fmt.Println(validateContainRank(s2, s1))

	s1 = "ab"
	s2 = "eidboaoo"
	fmt.Println(validateContainRank(s2, s1))
}
