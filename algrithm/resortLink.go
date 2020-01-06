package algrithm

//对链表进行重排序
// 1-2-3-4-5-6
//重排序之后为1-6-2-5-3-4
func ResortLink(root *Node) {
	if root == nil {
		return
	}
	if root.Next != nil && root.Next.Next == nil {
		return
	}
	slow, fast := root, root
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	cur := slow.Next
	slow.Next = nil
	var pre *Node
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	newroot := pre

	//newroot和root进行合并
	node1 := root
	node2 := newroot
	for node1 != nil && node2 != nil {
		tmp := node1.Next
		node1.Next = node2
		node1 = tmp
		tmp = node2.Next
		node2.Next = node1
		node2 = tmp
	}
}

func TestResortLink() {
	arr := []int{1,2,3,4,5, 6, 7, 8}
	root := GenerateLink(arr)
	ResortLink(root)

	PrintNodes(root)
}