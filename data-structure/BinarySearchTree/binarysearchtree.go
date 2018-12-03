package binarysearchtree

type node struct {
	data  int
	left  *node
	right *node
}

func NewNode(root int) *node {
	return &node{
		data: root,
	}
}

func (n *node) AddNode(data int) {
	if data < n.data {
		if n.left == nil {
			n.left = NewNode(data)
			return
		}
		n.left.AddNode(data)
	} else {
		if n.right == nil {
			n.right = NewNode(data)
			return
		}
		n.right.AddNode(data)
	}
	return
}

func DFSInOrder(root *node) (result []int) { // left -> root -> right
	tmp := root
	if tmp.left != nil {
		result = append(result, DFSInOrder(tmp.left)...)
	}
	result = append(result, tmp.data)
	if tmp.right != nil {
		result = append(result, DFSInOrder(tmp.right)...)
	}
	return result
}

func DFSPostOrder(root *node) (result string) { // right -> root -> left
	//tmp := root
	// using stack
	return result
}
