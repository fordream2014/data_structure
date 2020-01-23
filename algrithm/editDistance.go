package algrithm

import (
	"math"
	"fmt"
)

/*
给定两个单词 word1 和 word2，计算出将 word1 转换成 word2 所使用的最少操作数 。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符
示例 1:

输入: word1 = "horse", word2 = "ros"
输出: 3
解释:
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
 */

 //dist[i][j] 标示word1（长度为i）转换为 word2（长度为j）所使用的最少操作次数
 func getTransEditOperateNum(word1 , word2 []string) int {
	lenth1 := len(word1)
	lenth2 := len(word2)

	 dist := make([][]int, lenth1+1)
	 for i := 0; i < lenth1+1; i++ {
		 dist[i] = make([]int, lenth2+1)
	 }

	dist[0][0] = 0
	for j:=1; j<=lenth2; j++ {
		dist[0][j] = dist[0][j-1] + 1
	}
	 for i:=1; i<=lenth1; i++ {
		 dist[i][0] = dist[i-1][0] + 1
	 }

	 for i:=1; i<=lenth1; i++ {
	 	for j:=1; j<=lenth2; j++ {
			if word1[i-1] == word2[j-1] {
				dist[i][j] = dist[i-1][j-1]
			} else {
				min := math.Min(float64(dist[i][j-1]), float64(dist[i-1][j]))
				dist[i][j] = int(math.Min(float64(min), float64(dist[i-1][j-1]))) + 1
			}
		}
	 }
	 fmt.Println(dist)
 	return dist[lenth1-1][lenth2-1]
 }

 /*
 优化方案，空间复杂度由O(n*m)变为O(m)
  */
func optimization(word1 , word2 []string) int {
	lenth1 := len(word1)
	lenth2 := len(word2)

	dist := make([]int, lenth2+1)
	for j:=1; j<=lenth2; j++ {
		dist[j] = dist[j-1] + 1
	}

	var pre int
	for i:=1; i<=lenth1; i++ {
		for j:=1; j<=lenth2; j++ {
			if word1[i-1] == word2[j-1] {
				dist[j] = pre
			} else {
				min := math.Min(float64(dist[j-1]), float64(dist[j]))
				dist[j] = int(math.Min(float64(min), float64(pre))) + 1
				pre = dist[j]
			}
		}
	}
	fmt.Println(dist)
	return dist[lenth2-1]
}

 func TestMinDistance() {
 	word1 := []string{"h", "o", "r", "s", "3"}
 	word2 := []string{"r", "o", "s"}

 	//word1 := []string{"i","n","t","e","n","t","i","o","n"}
 	//word2 := []string{"e","x","e","c","u","t","i","o","n"}
 	dist := getTransEditOperateNum(word1, word2)
 	fmt.Println(dist)
 }
