package algrithm

import "fmt"

func validateNumeric(num string) bool {
	//数字的形式为 [+/-A].[B][E/e][C]
	//有小数点的话，A或者B不为空
	//有E/e的话，前面的数字一定不能为空，且后面的数字也不能为空
	index, is_numeric := scanNumer(num, 0)
	//fmt.Println(index, is_numeric)
	if index < len(num) && num[index] == '.' {
		num_index, temp := scanUnsignNumber(num, index+1)
		index = num_index
		is_numeric = is_numeric || temp
	}
	if index < len(num) && (num[index] == 'e' || num[index] == 'E') {
		if is_numeric {
			index, is_numeric = scanNumer(num, index+1)
		}
	}
	return is_numeric && index == len(num)
}

//无符号数字
func scanUnsignNumber(num string, index int) (int, bool) {
	start := index
	for index < len(num) {
		t := num[index]
		if t >= '0' && t <= '9' {
			index++
		} else {
			break
		}
	}
	is_numeric := false
	if index > start {
		is_numeric = true
	}
	return index, is_numeric
}

//有符号数字
func scanNumer(num string, index int) (int, bool) {
	if index >= len(num) {
		return index, false
	}
	if num[index] == '+' || num[index] == '-' {
		index++
	}
	return scanUnsignNumber(num, index)
}

//表示数值的字符串
//请实现一个函数用来判断字符串是否表示数值（包括整数和小数）
func TestNumeric() {
	//num := "123.45e+6"
	//fmt.Println(num, validateNumeric(num))
	//
	//num = "-123"
	//fmt.Println(num, validateNumeric(num))
	//
	//num = "3.1445"
	//fmt.Println(num, validateNumeric(num))
	//
	//num = "+100"
	//fmt.Println(num, validateNumeric(num))

	num := "12e"
	fmt.Println(num, validateNumeric(num))

	num = "12e+5.4"
	fmt.Println(num, validateNumeric(num))
}
