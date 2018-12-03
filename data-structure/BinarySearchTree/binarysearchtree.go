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
	// TODO: fix
	tmp := n
	if data < tmp.data {
		if tmp.left == nil {
			tmp.left = NewNode(data)

			return
		}
		tmp.left.AddNode(data)
	} else {
		if tmp.right == nil {
			tmp.right = NewNode(data)
			return
		}
		tmp.right.AddNode(data)
	}
	return
}

func DFSInOrder(root *node) (result []int) { // left -> root -> right
	tmp := root
	if tmp.left != nil {
		result = append(result, DFSInOrder(tmp.left)...)
	}
	// root
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
