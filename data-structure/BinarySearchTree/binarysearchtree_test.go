package binarysearchtree

import (
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	var root = NewNode(50)
	root.AddNode(30)
	root.AddNode(20)
	root.AddNode(40)
	root.AddNode(70)
	root.AddNode(60)
	root.AddNode(80)
	data := DFSInOrder(root)
	result := []int{20, 30, 40, 50, 60, 70, 80}
	for i := range data {
		if data[i] != result[i] {
			t.FailNow()
		}
	}
}
