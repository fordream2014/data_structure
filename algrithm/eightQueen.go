package algrithm

import "fmt"

//8皇后问题，同行、同列、对角线上只有一个
var result [8]int

func EightQueen() {
	setPoint(0)
}

func setPoint(row int) {
	if row == 8 {
		fmt.Println(result)
		return
	}
	for i:=0; i<8; i++ {
		if validateIsOk(row,i) { //剪枝
			result[row] = i
			setPoint(row+1)
		}
	}
}

//验证当前行、列上的点是否ok
func validateIsOk(row, column int) bool {
	if row == 0 {
		return true
	}
	var left = column - 1
	var right = column + 1
	for i:=row-1; i>=0; i-- {
		if result[i] == column {
			return false
		}
		if left >= 0 && result[i] == left {
			return false
		}
		if right < 8 && result[i] == right {
			return false
		}
		left--
		right++
	}
	return true
}