package algrithm

import (
	"strconv"
	"fmt"
)

func stringMultiplication(num1, num2 string) string {
	if validateEmpty(num1) || validateEmpty(num2) {
		return "0"
	}

	//num1Len := len(num1)
	num2Len := len(num2)
	var sums []string
	for i:=num2Len-1; i>=0; i-- {
		itemsum := getMultiplication(num1, string(num2[i]), num2Len-i-1)
		sums = append(sums, itemsum)
		//fmt.Println(sums)
	}
	sum := getSummary(sums)
	return sum
}

func getSummary(nums []string) string {
	if len(nums) == 0 {
		return "0"
	}
	if len(nums) == 1 {
		return nums[0]
	}

	sum := nums[0]
	for i:=1; i<len(nums); i++ {
		sum = sumTwoNum(sum, nums[i])
	}
	return sum
}

func sumTwoNum(num1, num2 string) string {
	fmt.Println(num1, num2)
	i := len(num1)-1
	j := len(num2)-1
	var result string
	carry := 0
	for i>=0 || j>=0 || carry > 0 {
		num1_t := 0
		if i >= 0 {
			num1_t = int(num1[i]-'0')
			i--
		}
		num2_t := 0
		if j >= 0 {
			num2_t = int(num2[j]-'0')
			j--
		}
		sum := num1_t + num2_t + carry
		carry = 0
		if sum >= 10 {
			sum -= 10
			carry = 1
		}
		result = strconv.Itoa(sum) + result
	}

	return result
}

//字符串相乘
func getMultiplication(num1, num2 string, zerofil int) string {
	var arr []int
	carry := 0  //进位

	for i:=0; i<zerofil; i++ {
		arr = append(arr, 0)
	}
	num2_int,_ := strconv.Atoi(num2)
	//fmt.Println(num2_int, num1)
	for i:=len(num1)-1; i>=0; i-- {
		up := int(num1[i] - '0')
		//fmt.Println(up)
		multi := up * num2_int + carry
		carry = 0
		if multi >= 10 {
			carry = multi/10
			multi = multi%10
		}
		arr = append(arr, multi)
	}
	if carry > 0 {
		arr = append(arr, carry)
	}
	//fmt.Println(arr)
	result := ""
	for i:=len(arr)-1; i>=0; i-- {
		result += strconv.Itoa(arr[i])
	}
	return result
}

//判读是否为0
func validateEmpty(num string) bool {
	return num == "" || num == "0"
}

/*
给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

示例 1:

输入: num1 = "2", num2 = "3"
输出: "6"
示例 2:

输入: num1 = "123", num2 = "456"
输出: "56088"
说明：

num1 和 num2 的长度小于110。
num1 和 num2 只包含数字 0-9。
num1 和 num2 均不以零开头，除非是数字 0 本身。
不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
 */
func TestStringMultiplication() {
	//num1 := "123"
	//num2 := "356"
	num1 := "2"
	num2 := "3"
	result := stringMultiplication(num1, num2)
	fmt.Println(result)
}
