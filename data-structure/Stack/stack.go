package Stack

type Node struct {
	data int
	prev *Node
}

type StackElm struct {
	Data Node
	top  *Node
}

func NewStack() *StackElm {
	return &StackElm{
		top: nil,
	}
}

func (se *StackElm) Push(data Node) {
	if se.top == nil {
		se.top = &data
		return
	}
	data.prev = se.top
	se.top = &data
	return
}

func (se StackElm) Peek() Node {
	return *se.top
}
func (se *StackElm) Pop() Node {
	data := se.top
	se.top = se.top.prev
	return *data
}

func (se *StackElm) isEmpty() bool {
	if se.top == nil {
		return true
	}
	return false
}
