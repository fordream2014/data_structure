package algrithm

import (
	"math"
	"fmt"
)

//0-1背包问题
// 现有n个物品，每个物品的重量不等，背包总承重W kg，期望选择几件物品，
// 装载到背包中，在不超过背包所能装载重量的前提下，如何让背包中物品的总重量最大
var max_weight int = math.MinInt64
var package_limit = 10

func setItem(arr []int, index, current_weight int) {
	if current_weight == package_limit || index >= len(arr){
		if current_weight > max_weight {
			max_weight = current_weight
		}
		return
	}

	setItem(arr, index+1, current_weight)
	if current_weight + arr[index] <= package_limit {
		setItem(arr, index+1, current_weight + arr[index])
	}
}

func Test01Package() {
	arr := []int{1,2,2,3,4,2,4}
	setItem(arr, 0, 0)
	fmt.Printf("0-1背包问题，最大重量为%v", max_weight)
}
