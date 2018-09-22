package LinkedList

type TData int

type ElementDbl struct {
	Data TData

	next *ElementDbl
	prev *ElementDbl
}

type DoublyLinkedList struct {
	Head *ElementDbl
	Tail *ElementDbl
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		Head: nil,
		Tail: nil,
	}
}

func (e *DoublyLinkedList) AddFirst(elm ElementDbl) {
	newElement := elm
	if e.Head == nil {
		e.Head = &newElement
		e.Tail = e.Head
		return
	}
	newElement.next = e.Head
	e.Head.prev = &newElement
	e.Head = &newElement
	return
}
func (e *DoublyLinkedList) AddLast(elm ElementDbl) {
	newDbl := elm
	if e.Head == nil {
		e.Head = &newDbl
		e.Tail = e.Head
	}
	newDbl.prev = e.Tail
	e.Tail.next = &newDbl
	e.Tail = &newDbl
	return
}

func (e *DoublyLinkedList) DeleteFirst() {

	e.Head = e.Head.next
	if e.Head == nil {
		e.Tail = nil
	}

	return
}

func (e *DoublyLinkedList) DeleteAtPostion(index int) {
	i := 0
	tmp := e.Head
	for tmp != e.Tail {
		if i == index {
			tmp.prev.next = tmp.next
			tmp.next.prev = tmp.prev
			return
		}
		i++
		tmp = tmp.next
	}
}

func (e *DoublyLinkedList) FindElement(elm TData) int {
	i := 0
	tmp := e.Head
	for tmp != e.Tail {
		if tmp.Data == elm {
			return i
		}
		i++
		tmp = tmp.next
	}
	return -1
}
