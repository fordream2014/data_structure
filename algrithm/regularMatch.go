package algrithm

import "fmt"

func regularMatch(str, regular string) bool {

	return matchStr(str, regular, 0, 0)
}

func matchStr(str, regular string, str_ind, reg_ind int) bool {
	if reg_ind >= len(regular) {
		//正则结束
		if str_ind >= len(str) {
			return true
		} else {
			return false
		}
	}

	if reg_ind + 1 < len(regular) && regular[reg_ind+1] == '*' {
		if str_ind < len(str) && (regular[reg_ind] == '.' || regular[reg_ind] == str[str_ind]) {
			return matchStr(str, regular, str_ind + 1, reg_ind + 2) ||
				matchStr(str, regular, str_ind + 1, reg_ind) ||
				matchStr(str, regular, str_ind, reg_ind + 2)
		} else {
			return matchStr(str, regular, str_ind, reg_ind + 2)
		}
	} else {
		//fmt.Println(reg_ind, str_ind)
		if str_ind < len(str) && (regular[reg_ind] == '.' || regular[reg_ind] == str[str_ind]) {
			return matchStr(str, regular, str_ind + 1, reg_ind + 1)
		}
	}

	return false
}

//题目介绍：实现一个函数用来匹配包含'.'和'*'的正则表达式。模式中的'.'表示任意一个字符,而'*'表示它前面的字符可以出现任意次.
//例如字符串'aaa'与模式'a.a'和'ab*ac*a'匹配，但与'aa.a'和'ab*a'均不匹配
func TestRegular() {
	str := "aaa"
	regular := "ab*ac*a"
	res := regularMatch(str, regular)
	fmt.Println(res)
}
