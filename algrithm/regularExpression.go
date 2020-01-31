package algrithm

import "fmt"

// "*"匹配任意多个（大于等于0个）任意字符
// "?" 匹配零个或者一个任意字符
var match bool

func matchRegularExpression(text, regular []string, text_i, regular_i int) {
	if match {
		return
	}
	if regular_i == len(regular) {
		if text_i == len(text) {
			match = true
		}
		return
	}

	if regular[regular_i] == "*" {
		for i:=text_i+1; i<len(text); i++ {
			matchRegularExpression(text, regular, i, regular_i+1)
		}
	} else if regular[regular_i] == "?" {
		matchRegularExpression(text, regular, text_i+1, regular_i+1)
		matchRegularExpression(text, regular, text_i, regular_i+1)
	} else if regular[regular_i] == text[text_i]{
		matchRegularExpression(text, regular, text_i+1, regular_i+1)
	}
}

func TestMatchRegularExpression() {
	text := []string{"a","c", "1", "2"}
	regular := []string{"a", "*", "2"}

	matchRegularExpression(text, regular, 0, 0)
	fmt.Printf("正则表达式匹配结果: %v", match)
}


