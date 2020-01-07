package algrithm

//相邻元素实现反转，例如: 1-2-3-4-5-6-7
//输出：2-1-4-3-6-5-7
func inverseAdjacentElem(root *Node) *Node {
	if root == nil {
		return nil
	}

	newroot := &Node{
		Val: -1,
		Next: root,
	}
	pre := newroot
	cur := root
	for cur != nil && cur.Next != nil {
		next := cur.Next.Next
		pre.Next = cur.Next
		cur.Next.Next = cur
		cur.Next = next
		pre = cur
		cur = next
	}
	return newroot.Next
}

func TestInverseAdjacentElem() {
	arr := []int{1,2,3,4,5,6,7}
	root := GenerateLink(arr)

	newroot := inverseAdjacentElem(root)
	PrintNodes(newroot)
}
