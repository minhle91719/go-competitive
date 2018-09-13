package LinkedList

import "fmt"

type tData int

type Elememt struct {
	data tData
	next *Elememt
}

type SimpleLinkedList struct {
	Header *Elememt
}

// NewSimpleLinkedList create instance for linked_list
func NewSimpleLinkedList() *SimpleLinkedList {
	return &SimpleLinkedList{Header: nil}
}

func (sll *SimpleLinkedList) AddElement(data tData) {
	if sll.Header == nil {
		sll.Header = &Elememt{data: data, next: nil}
		return
	}
	//	fmt.Println(data)
	nElm := Elememt{data: data, next: sll.Header}
	sll.Header = &nElm
}

func (sll *SimpleLinkedList) DeleteElement(data tData) {
	var previous *Elememt

	tmp := sll.Header

	for tmp != nil {
		if tmp.data == data {
			previous.next = tmp.next
			return
		}
		previous = tmp
		tmp = tmp.next
	}
}

func (sll SimpleLinkedList) ExistData(data tData) bool {
	tmp := sll.Header
	for tmp != nil {
		if tmp.data == data {
			return true
		}
		tmp = tmp.next
	}
	return false
}

// sort SimpleLinkedList

func (sll *SimpleLinkedList) Sort(asc bool) {
	tmpHeader := sll.Header
	for tmpHeader != nil {
		tmpH2 := tmpHeader.next
		for tmpH2 != nil {
			if asc {
				if tmpHeader.data > tmpH2.data {
					tmpHeader.data, tmpH2.data = tmpH2.data, tmpHeader.data
				}
				continue
			}
			if tmpHeader.data < tmpH2.data {
				tmpHeader.data, tmpH2.data = tmpH2.data, tmpHeader.data
			}
		}
	}
	return
}

// [ 5 4 3 2 1] => [1 2 3 4 5]

func (sll *SimpleLinkedList) Reverse() {
	var (
		current          = sll.Header
		prev    *Elememt = nil
	//	ne      *Elememt = nil
	)

	for current != nil {
		if prev == nil {
			prev = current
			current = current.next
			prev.next = nil
			continue
		}

		tmp := current.next
		if current.next == nil {
			current.next = prev
			sll.Header = current
			return
		}

		current.next = prev
		prev = current
		current = tmp
	}
}

func (sll SimpleLinkedList) Show() {
	tmpH := sll.Header
	for tmpH != nil {
		fmt.Println(tmpH.data)
		tmpH = tmpH.next
	}
}

func (sll SimpleLinkedList) ToArray() []tData {
	var (
		arr []tData
	)
	tmpH := sll.Header
	for tmpH != nil {
		arr = append(arr, tmpH.data)
		tmpH = tmpH.next
	}
	return arr
}
