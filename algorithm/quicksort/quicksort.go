package quicksort

func QuickSort(arr []int, start, end int) {
	if start < end {
		pivot := arr[end]
		i := start - 1
		for j := start; j <= end-1; j++ {
			if arr[j] <= pivot {
				i++
				swap(arr, j, i)
			}
		}
		swap(arr, end, i+1)
		QuickSort(arr, start, i)
		QuickSort(arr, i+2, end)
	}
}
func swap(arr []int, x int, y int) {
	tmp := arr[x]
	arr[x] = arr[y]
	arr[y] = tmp
}
