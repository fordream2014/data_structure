package algrithm

import "fmt"

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

func TestRevertKElem() {
	k := 3
	arr := []int{1,2,3,4,5,6,7,8,9}
	root := GenerateLink(arr)
	PrintNodes(root)
	fmt.Println()

	newroot := reverseKElem(root, k)
	PrintNodes(newroot)
}