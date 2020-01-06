package algrithm

import "fmt"

type BinaryTreeNode struct {
	Val int
	Left *BinaryTreeNode
	Right *BinaryTreeNode
}

type SpecialNode struct {
	Val string
	Left *SpecialNode
	Right *SpecialNode
	Parent *SpecialNode
}

//前序遍历
func PrintPreOrderList(root *BinaryTreeNode) {
	if root == nil {
		return
	}

	fmt.Printf("%v ", root.Val)
	if root.Left != nil {
		PrintPreOrderList(root.Left)
	}
	if root.Right != nil {
		PrintPreOrderList(root.Right)
	}
}

//中序遍历
func PrintMidOrderList(root *BinaryTreeNode) {
	if root == nil {
		return
	}

	if root.Left != nil {
		PrintMidOrderList(root.Left)
	}
	fmt.Printf("%v ", root.Val)
	if root.Right != nil {
		PrintMidOrderList(root.Right)
	}
}

type Node struct {
	Val int
	Next *Node
}

func PrintNodes(root *Node) {
	cur := root
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}

