package LinkedList

import (
	"testing"
)

func TestReverse(t *testing.T) {

	ArrayInt := []tData{5, 4, 3, 2, 1, 0}

	lkl := NewSimpleLinkedList()
	for _, v := range ArrayInt {
		lkl.AddElement(v)
	}
	// lkl.Show()
	// lkl.Reverse()
	// fmt.Println()
	// lkl.Show()
	lkl.Reverse()
	arraynew := lkl.ToArray()
	totalelm := len(arraynew) - 1
	for i, v := range ArrayInt {
		index := totalelm - (totalelm - i)
		sosanh := arraynew[index]
		if v != sosanh {
			t.Errorf("%s %d %d", "Khong khop: ", v, sosanh)
		}
	}

}
