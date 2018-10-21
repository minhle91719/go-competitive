package quicksort

import "testing"

func TestQuickSort(t *testing.T) {
	type args struct {
		arr   []int
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				arr:   []int{1, 2, 6, 8, 9, 6, 3, 2, 1, 4, 5, 5, 6, 9, 5, 8, 7, 4, 1, 2, 5, 6, 3, 2},
				start: 0,
				end:   23,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort(tt.args.arr, tt.args.start, tt.args.end)
			for i := 0; i < len(tt.args.arr)-1; i++ {
				if tt.args.arr[i] > tt.args.arr[i+1] {
					t.Error("fail test: :" + tt.name)
				}
			}
		})
	}
}
