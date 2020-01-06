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
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
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

func TestDeleteDup() {
	arr := []int{1,2,2,2,3,3,5}
	root := GenerateLink(arr)
	PrintNodes(root)
	fmt.Println()

	root = DeleteDup(root)
	PrintNodes(root)

}

