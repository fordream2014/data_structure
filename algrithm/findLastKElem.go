package algrithm

import "fmt"

//使用快慢指针，找到倒数第K个元素
func findLastKElem(root *Node, k int) *Node {
	if root == nil {
		return nil
	}
	if k <= 0 {
		return nil
	}

	fast := root
	num := 1
	for fast != nil {
		if num == k {
			break
		}
		fast = fast.Next
		num ++
	}

	if num < k {
		return nil
	}

	slow := root
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

func TestFindLastKElem() {
	//arr := []int{1,2,3,4,5,6}
	arr := []int{1,2,3,4,5,6,7}
	root := GenerateLink(arr)
	node := findLastKElem(root, 3)
	fmt.Println(node)
}
