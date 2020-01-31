package algrithm

import "fmt"

//删除有序链表中的所有重复项
//1-2-2-2-3-3-5  => 1-5
func DeleteDup(root *Node) *Node {
	if root == nil {
		return nil
	}
	if root.Next == nil {
		return root
	}
	newroot := &Node{
		Val: -1,
		Next: root,
	}
	pre := newroot
	cur := root
	for cur != nil {
		if cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		} else {
			if pre.Next == cur {
				pre = cur
				cur = cur.Next
			} else {
				pre.Next = cur.Next
				cur = cur.Next
			}
		}
	}
	return newroot.Next
}

//删除有序链表中的重复项
//1-2-2-2-3-3-5  => 1-2-3-5
//可以认为cur.Next == nil的情况和cur.Next.val == cur.Val的情况相同
func DeleteDup2(root *Node) *Node {
	if root == nil {
		return nil
	}
	if root.Next == nil {
		return root
	}

	pre := root
	cur := root.Next
	for cur != nil {
		if cur.Val == pre.Val {
			cur = cur.Next
		} else {
			if pre.Next != cur {
				pre.Next = cur

			}

			pre = cur
			cur = cur.Next
		}
	}
	return root
}

func TestDeleteDup() {
	arr := []int{1,2,2,2,3,3,5}
	root := GenerateLink(arr)
	PrintNodes(root)
	fmt.Println()

	root = DeleteDup(root)
	PrintNodes(root)

	fmt.Println()

	//方法2 1-2-3-5
	//arr = []int{1,2,2,2,3,3,5}
	//root = GenerateLink(arr)
	//root = DeleteDup2(root)
	//PrintNodes(root)
}

