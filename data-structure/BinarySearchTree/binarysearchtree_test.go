package binarysearchtree

import (
	"fmt"
	"testing"
)

func TestAddNode(t *testing.T) {
	// type args struct {
	// 	data int
	// }
	// tests := []struct {
	// 	name string
	// 	n    *node
	// 	args args
	// }{
	// 	// TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		tt.n.AddNode(tt.args.data)
	// 	})
	// }

	var root = &node{
		data: 1,
		left: &node{
			data: 7,
			left: &node{
				data: 2,
			},
			right: &node{
				data: 6,
			},
		},
		right: &node{
			data: 8,
		},
	}
	// var listleaf = []int{2, 7, 6, 8}
	// for _, v := range listleaf {
	// 	root.AddNode(v)
	// }
	fmt.Println(DFSInOrder(root))
}
