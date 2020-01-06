package algrithm

import "fmt"

//使用链表存储字符串，判断是否是回文
//使用快慢指针，同时慢指针反转链表
func validateHuiwen(root *Node) bool {

	if root == nil {
		return false
	}
	if root.Next == nil {
		return true
	}

	var cur, pre *Node
	slow, fast := root, root
	cur = slow
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		//修改指针方向
		cur.Next = pre
		pre = cur
		cur = slow
	}

	//可以借助画图，得到下列代码
	if fast.Next == nil {
		//为奇数个
		slow = slow.Next
		cur = pre
	} else {
		slow = slow.Next
		cur.Next = pre
	}

	for cur != nil {
		if cur.Val != slow.Val {
			return false
		}
		//这一步很重要
		cur = cur.Next
		slow = slow.Next
	}
	return true
}

func generateHuiwen(arr []int) *Node {
	if len(arr) == 0 {
		return nil
	}
	root := &Node {
		Val: -1,
	}
	cur := root
	for _,v := range arr {
		cur.Next = &Node{
			Val: v,
		}
		cur = cur.Next
	}
	return root.Next
}

func TestValidateHuiwen() {

	arr := []int{1,2,3,3,2,1}
	root := generateHuiwen(arr)

	huiwen := validateHuiwen(root)
	fmt.Println(huiwen)

	arr = []int{1,2,3,4,3,2,1}
	root = generateHuiwen(arr)
	huiwen = validateHuiwen(root)
	fmt.Println(huiwen)
}

