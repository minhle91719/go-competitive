package linkedlist

type Node struct {
	data interface{}
	next *Node
}

func (n *Node) setNext(nNext *Node) {
	n.next = nNext
}

type simplyLinkedList struct {
	head *Node
}
type SimplyLinkedList interface {
	AddNodeWithValue(value interface{})
	AddNode(node *Node)
	RemoveValueLinkedList(value string) 

	ToArray() []interface{}
}
