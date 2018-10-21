package quicksort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	var newarr = []int{1, 2, 3, 4, 8, 9, 2, 3, 4, 5, 6}
	QuickSort(newarr, 0, len(newarr)-1)

	for _, v := range newarr {
		fmt.Println(v)
	}
}
