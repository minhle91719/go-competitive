package binarysearchtree

import "fmt"

type node struct {
	data  int
	left  *node
	right *node
}

func NewTree(root int) *node {
	return &node{
		data: root,
	}
}

func (n *node) AddNode(data int) {
	// TODO: fix
	tmp := n
	for tmp != nil {
		if tmp.data > data { // left
			if tmp.left == nil {
				tmp.left = &node{
					data: data,
				}
				return
			}
			tmp = tmp.left
			continue
		}
		if tmp.data < data { // right
			if tmp.right == nil {
				tmp.right = &node{
					data: data,
				}
				return
			}
			tmp = tmp.right
		}
		if tmp.data == data {
			return
		}
	}
}

func DFSInOrder(root *node) (result string) { // left -> root -> right
	tmp := root
	if tmp.left != nil {
		result += DFSInOrder(tmp.left)
	}
	// root
	result += fmt.Sprintf("%d ", tmp.data)
	if tmp.right != nil {
		result += DFSInOrder(tmp.right)
	}
	return result
}

func DFSPostOrder(root *node) (result string) { // right -> root -> left
	//tmp := root
	// using stack
	return result
}
