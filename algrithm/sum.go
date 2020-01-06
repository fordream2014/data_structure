package algrithm

//计算两个单链表所代表的数之和
//有几点需要注意
//1. 指针的递进
//2. sum的归零
//3. 上进位的条件：>= 10
func sumLink(left, right *Node) *Node {
	root := &Node{
		Val: -1,
	}
	sum := 0
	cur := root
	for left != nil || right != nil {
		leftVal, rightVal := 0, 0
		if left != nil {
			leftVal = left.Val
			left = left.Next
		}
		if right != nil {
			rightVal = right.Val
			right = right.Next
		}
		val := leftVal + rightVal + sum
		sum = 0
		if val >= 10 {
			sum = val/10
			val = val % 10
		}

		node := &Node{
			Val: val,
		}
		cur.Next = node
		cur = cur.Next
	}
	if sum != 0 {
		node := &Node{
			Val: sum,
		}
		cur.Next = node
	}

	return root.Next
}

func TestSum() {
	left := []int{3,1,5}
	right := []int{5,9,2}
	leftroot := GenerateLink(left)
	rightroot := GenerateLink(right)

	sum := sumLink(leftroot, rightroot)
	PrintNodes(sum)
}
