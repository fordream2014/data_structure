package algrithm

//合并两个有序链表
func mergeLink(node1, node2 *Node) *Node {
	newroot := &Node{
		Val: -1,
	}
	pre := newroot
	for node1 != nil && node2 != nil {
		if node1.Val > node2.Val {
			pre.Next = node2
			pre = node2
			node2 = node2.Next
		} else {
			pre.Next = node1
			pre = node1
			node1 = node1.Next
		}
	}

	if node1 != nil {
		pre.Next = node1
	}
	if node2 != nil {
		pre.Next = node2
	}
	return newroot.Next
}

func TestMergeLink() {
	arr1 := []int{1,3,4,7,9}
	arr2 := []int{2,5,6,8}

	node1 := GenerateLink(arr1)
	node2 := GenerateLink(arr2)
	root := mergeLink(node1, node2)
	PrintNodes(root)
}
