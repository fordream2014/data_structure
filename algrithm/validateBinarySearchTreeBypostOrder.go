package algrithm

import "fmt"

//根据后序遍历判断是否是二叉搜索树
func validateBSTreeByPostOrder(str []int, left, right int) bool {
	if left >= right {
		return false
	}

	root := str[right]
	i := left
	for ; i<=right-1; i++ {
		if str[i] > root {
			break
		}
	}

	for j:=i; j<=right-1; j++ {
		if str[j] <= root {
			return false
		}
	}
	//fmt.Println("mid partition ind:", i, root)

	if i - 1 - left > 0 {
		left_res := validateBSTreeByPostOrder(str, left, i-1)
		if !left_res {
			return left_res
		}
	}
	if right - 1 - i > 0 {
		right_res := validateBSTreeByPostOrder(str, i, right-1)
		if !right_res {
			return right_res
		}
	}

	return true
}

func TestValidateBSTreeByPostOrder() {
	arr := []int{5,7,6,9,11,10,8}
	//arr := []int{7,4,6,5}
	res := validateBSTreeByPostOrder(arr, 0, len(arr)-1)
	fmt.Println(res)
}
