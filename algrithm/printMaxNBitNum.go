package algrithm

import "fmt"

func print_max_n_bit_number(length int) {
	if length == 0 {
		fmt.Println("no number exist")
		return
	}

	arr := make([]byte, length)
	fill_bit(arr, 0, length)
}

//使用全排列的方式
func fill_bit(arr []byte, index, max_len int) {
	if index == max_len {
		fmt.Println(string(arr))
		return
	}

	for i:=0; i<10; i++ {
		arr[index] = byte('0' + i)
		fill_bit(arr, index+1, max_len)
	}
}

//方法2，使用字符串存储大数
func print_max_n_bit_number2(length int) {
	num := make([]byte, length)
	for i:=0; i<length; i++ {
		num[i] = '0'
	}
	for !string_increase(num) {
		fmt.Println(string(num))
	}
}

//字符串递增操作
func string_increase(num []byte) bool{
	full := false
	max := len(num)
	extend := 0
	for i:=max-1; i>=0; i-- {
		var t int
		if i == max-1 {
			t = int(num[i] - '0') + 1
		} else {
			t = int(num[i] - '0') + extend
			extend = 0
		}
		if t >= 10 {
			if i == 1 {
				full = true
			}
			extend = 1
			t -= 10
			num[i] = byte(t) + '0'
		} else {
			num[i] = byte(t) + '0';
			break
		}
	}
	return full
}

//输入数字n，按顺序打印出从1到最大的n位十进制数
//需要注意的地方：大数的存储，使用int可能造成内存溢出
func TestPrintMaxNBitNumber() {
	num := 4
	print_max_n_bit_number2(num)
	fmt.Println("*****************************")
}
