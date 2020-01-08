package algrithm

import (
	"math"
	"fmt"
)

//求一个数的平方根
//牛顿迭代公式：x[n+1] = x[n] - f(x[n])/f'(x[n])
//以上公式会最终收敛到零点，x[n+1]最终会收敛到f(x)=0的零点
//对于一个数的平方根：即求解f(x)=x^2-a=0，迭代公式为：x[n+1]=(x[n]+a/x[n])/2

func sqrt(base float64, degree float64) float64 {
	tmp := base
	for math.Pow(tmp, 2) - base > degree {
		tmp = (tmp + base/tmp)/2
	}
	return tmp
}

func TestSqrt() {
	x := 2.0
	s := sqrt(x, 0.000001)
	fmt.Println(s)
}

