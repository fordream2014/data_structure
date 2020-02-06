package algrithm

type BSTreeNode struct {
	Val int
	Left *BSTreeNode
	Right *BSTreeNode
}

type BSTree struct {
	root *BSTreeNode
}

//二叉查找树插入节点
func (this *BSTree) add(data int) bool{
	if this.root == nil {
		this.root = &BSTreeNode{
			Val: data,
		}
		return true
	}

	p := this.root
	for p != nil {
		if p.Val < data {
			if p.Right == nil {
				p.Right = &BSTreeNode{
					Val:data,
				}
				return true
			} else {
				p = p.Right
			}
		} else {
			if p.Left == nil {
				p.Left = &BSTreeNode{
					Val: data,
				}
				return true
			} else {
				p = p.Left
			}
		}
	}
	return false
}

//二叉查找树的查找操作
func (this *BSTree) search(data int) *BSTreeNode {
	p := this.root
	for p != nil {
		if p.Val < data {
			p = p.Right
		} else if p.Val > data {
			p = p.Left
		} else {
			break
		}
	}

	return p
}

//删除节点，可以分为三种情况
//1. 删除叶节点
//2. 删除节点，有一个子节点

func (this *BSTree) delete(data int) bool {
	if this.root == nil {
		return false
	}

	var pp *BSTreeNode
	p := this.root
	for p != nil {
		if p.Val > data {
			pp = p
			p = p.Left
		} else if p.Val < data {
			pp = p
			p = p.Right
		} else {
			break
		}
	}

	if p == nil { //没有找到
		return false
	}

	if p.Left != nil && p.Right != nil {
		minpp := p
		minp := p.Right
		for minp.Left != nil {
			minpp = minp
			minp = minp.Left
		}

		p.Val = minp.Val
		p = minp
		pp = minpp
	}

	var child *BSTreeNode
	if p.Left != nil {
		child = p.Left
	} else if p.Right != nil {
		child = p.Right
	}

	if pp == nil {
		this.root = child
	} else if pp.Left == p {
		pp.Left = child
	} else {
		pp.Right = child
	}
	return true
}