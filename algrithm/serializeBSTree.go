package algrithm

import (
	"fmt"
	"strconv"
)

//使用前序遍历
//序列化
func serialize(node *BSTreeNode, stream string) string {
	if node == nil {
		stream = stream + "$"
		return stream
	}

	val := node.Val
	stream += strconv.Itoa(val)
	//fmt.Println(strconv.Itoa(val))
	stream = serialize(node.Left, stream)
	stream = serialize(node.Right, stream)
	return stream
}

var stream_g_ind int

func unserialize(stream string) *BSTreeNode {
	if stream_g_ind >= len(stream) {
		return nil
	}
	//fmt.Println(stream_g_ind)
	if stream[stream_g_ind] == '$' {
		return nil
	} else {
		node := &BSTreeNode{
			Val: int(stream[stream_g_ind] - '0'),
		}
		//fmt.Println(stream_g_ind, int(stream[stream_g_ind] - '0'))
		stream_g_ind = stream_g_ind + 1
		if stream_g_ind < len(stream) {
			node.Left = unserialize(stream)
		}

		stream_g_ind = stream_g_ind + 1
		if stream_g_ind < len(stream) {
			node.Right = unserialize(stream)
		}
		return node
	}
}

func TestSerializeBSTree() {
	root := &BSTreeNode{
		Val: 1,
	}
	head := root
	left := &BSTreeNode{
		Val: 2,
	}
	right := &BSTreeNode{
		Val: 5,
	}
	root.Left = left
	root.Right = right

	root = left
	left1 := &BSTreeNode{
		Val: 3,
	}
	right1 := &BSTreeNode{
		Val: 4,
	}
	root.Left = left1
	root.Right = right1

	root = right
	left2 := &BSTreeNode{
		Val: 6,
	}
	right2 := &BSTreeNode{
		Val: 7,
	}
	root.Left = left2
	root.Right = right2

	var stream string = ""
	stream = serialize(head, stream)

	fmt.Println("序列化的结果", stream)

	newhead := unserialize(stream)
	fmt.Println(newhead.Left.Right, newhead.Right.Right)
}
