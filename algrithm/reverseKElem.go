package algrithm

import (
	"fmt"
	"math"
)

//如何把链表以K个节点为一组进行反转
//1 先找到k的节点
//2 设置最后一个节点的Next=nil
//3 反转k个节点
//4 k个节点前后两个指针设置
func reverseKElem(root *Node, k int) *Node {
	if root == nil {
		return nil
	}

	newroot := &Node{
		Val: -1,
		Next: root,
	}
	pre := newroot
	begin := root
	for begin != nil {
		end := begin
		for i:=1;i<k;i++ {
			if end.Next != nil {
				end = end.Next
			} else {
				return newroot.Next
			}
		}

		next := end.Next
		end.Next = nil
		pre.Next = revertSubLink(begin)
		begin.Next = next

		pre = begin
		begin = next
	}
	return newroot.Next
}

//反转k个元素
func revertSubLink(node *Node) *Node {
	var pre, cur *Node
	cur = node
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}

	return pre
}

//开头不足k个的，不反转
//例如： {1,2,3,4,5,6,7} k=2
//反转结果为：{1,3,2,5,4,7,6}
func reverseKElemVariety(root *Node, k int) *Node {
	if root == nil || k <= 1 {
		return root
	}
	length := getLinkListLen(root)
	if length < k {
		return root
	}
	mod := length%k

	newroot := &Node{
		Val: math.MaxInt64,
		Next: root,
	}
	pre := newroot
	cur := root
	for num:=0; num<mod;num++ {
		pre = cur
		cur = cur.Next
	}

	for cur != nil {
		pre.Next = reverseSubLinkVariety(cur, k)
		pre = cur
		cur = cur.Next
	}
	return newroot.Next
}

func reverseSubLinkVariety(node *Node, k int) *Node {
	var pre *Node
	cur := node
	//o-o-o
	for num:=0; num<k; num++ {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	node.Next = cur
	return pre
}

//获取链表长度
func getLinkListLen(root *Node) int {
	if root == nil {
		return 0
	}
	cur := root
	len := 0
	for cur != nil {
		cur = cur.Next
		len++
	}
	return len
}

func TestRevertKElem() {
	k := 2
	arr := []int{1,2,3,4,5,6,7,8,9}
	root := GenerateLink(arr)
	PrintNodes(root)
	fmt.Println()

	//newroot := reverseKElem(root, k)
	//PrintNodes(newroot)

	newroot := reverseKElemVariety(root, k)
	PrintNodes(newroot)
}