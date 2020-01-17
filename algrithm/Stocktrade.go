package algrithm

import "fmt"

/*
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

示例 1:

输入: [7,1,5,3,6,4]
输出: 7
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
     随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。
 */
//暴力法求解
func StockTradeByForce(prices []int, s int) int {
	if s >= len(prices) - 1 {
		return 0
	}

	max := 0
	for start:=s; start<len(prices); start++ {
		maxprofile := 0
		for i:=start+1;i<len(prices);i++ {
			if prices[i] > prices[start] {
				profile := prices[i] - prices[start] + StockTradeByForce(prices, i+1)

				if profile > maxprofile {
					maxprofile = profile
				}
			}
		}

		if maxprofile > max {
			max = maxprofile
		}
	}

	return max
}

//贪心算法求解
func StockTradeByGreedy(prices []int) int {
	max := 0

	for i:=1; i<len(prices); i++ {
		if prices[i] > prices[i-1] {
			max += prices[i] - prices[i-1]
		}
	}
	return max
}

func TestStockTrade() {
	//arr := []int{7,1,5,3,6,4}
	//arr := []int{1,2,3,4,5}
	//arr := []int{7,6,4,3,1}
	//max := StockTradeByForce(arr, 0)
	//fmt.Printf("暴力法求解股票问题，结果%v \n", max)

	//arr := []int{7,1,5,3,6,4}
	arr := []int{1,2,3,4,5}
	//arr := []int{7,6,4,3,1}
	max := StockTradeByGreedy(arr)
	fmt.Printf("贪心法求解股票问题，结果%v \n", max)
}
