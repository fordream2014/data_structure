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
func TestStringRank() {
	str := "abcb"
	stringRank(str)

	fmt.Println(ranks)
}
