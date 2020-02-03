package algrithm

import "fmt"

//保存所有的排列组合
var ranks []string

func stringRank(base string) []string {
	if len(base) <= 1 {
		return []string{base}
	}
	var chars = []byte(base)
	substringRank(chars, 0)
	return ranks
}

func substringRank(base []byte, begin int) {
	if begin == len(base) {
		exist := false
		base_str := string(base)
		for _,v := range ranks {
			if v == base_str {
				exist = true
				break
			}
		}
		if !exist {
			ranks = append(ranks, base_str)
		}
		return
	}

	for i:=begin; i<len(base); i++ {
		temp := base[begin]
		base[begin] = base[i]
		base[i] = temp

		substringRank(base, begin+1)

		temp = base[i]
		base[i] = base[begin]
		base[begin] = temp
	}
}

//输入一个字符串，打印出该字符串中字符的所有排列，例如：
//输入abc，则打印出由字符a,b,c所能排列出来的所有字符串：abc,acb,bac,bca,cab,cba

//注意点：输入的字符串可能由重复的字符
//在每个分支进行完后，要将交换过的元素复位，从而不会影响其他分支。
//因为有“Swap A with A”这样的操作，并且题目指出可能有字符重复，所以分支末尾可能有重复的序列，在加入ArrayList要进行去重判断，不重复再加入。

//【解题思路】
//把字符串分为两部分：一部分是字符串的第一个字符，另一部分是第一个字符以后的所有字符。
//
//第一步是求所有可能出现在第一个位置的字符，即把第一个字符和后面所有字符交换。（for循环、交换操作）
//第二步是固定住第一个字符，求后面所有字符的排列。（递归）
//而“求后面所有字符的排列”即可按照上面的思路递归进行。

func TestStringRank() {
	str := "abcb"
	stringRank(str)

	fmt.Println(ranks)
}
