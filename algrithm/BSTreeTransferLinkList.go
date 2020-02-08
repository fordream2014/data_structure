package algrithm

import "fmt"

func bstreeTransfer(root *BSTreeNode) *BSTreeNode {
	if root == nil {
		return root
	}

	var last *BSTreeNode
	last = convertBSTree(root, last)

	//返回头指针
	head := last
	for head.Left != nil {
		head = head.Left
	}
	return head
}

func convertBSTree(cur, last *BSTreeNode) *BSTreeNode {
	if cur == nil {
		return nil
	}
	if cur.Left != nil {
		last = convertBSTree(cur.Left, last)
	}

	cur.Left = last
	if last != nil {
		last.Right = cur
	}

	last = cur
	if cur.Right != nil {
		last = convertBSTree(cur.Right, last)
	}
	return last
}

//题目要求：输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的双向链表。
// 要求不能创建任何新的节点，只能调整树中节点指针的指向
func TestBSTreeConvert() {
	root := &BSTreeNode{
		Val: 10,
	}
	head := root
	left := &BSTreeNode{
		Val: 6,
	}
	right := &BSTreeNode{
		Val: 14,
	}
	root.Left = left
	root.Right = right

	root = left
	left1 := &BSTreeNode{
		Val: 4,
	}
	right1 := &BSTreeNode{
		Val: 8,
	}
	root.Left = left1
	root.Right = right1

	root = right
	left2 := &BSTreeNode{
		Val: 12,
	}
	right2 := &BSTreeNode{
		Val: 16,
	}
	root.Left = left2
	root.Right = right2

	head = bstreeTransfer(head)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Right
	}
}
