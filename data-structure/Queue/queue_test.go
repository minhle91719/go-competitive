package queue

import (
	"fmt"
	"log"
	"testing"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue(20)
	q.Enqueue(8)
	q.Enqueue(6)
	q.Enqueue(7)
	fmt.Println(q.Capacity(), q.Length(), q.ToSliceWithDequeue())
	a, _ := q.Dequeue()
	b, _ := q.Dequeue()
	c, _ := q.Dequeue()
	_, err := q.Dequeue()
	if err != nil {
		log.Println(err)
	}

	fmt.Println(a, b, c)
}

func Test_queue_Length(t *testing.T) {
	queue5Value := NewQueue(5)
	queue5Value.EnqueueWithSlice([]interface{}{1, 2, 3, 4, 5})
	tests := []struct {
		name   string
		fields IQueue
		want   uint32
	}{
		{
			name:   "5",
			fields: queue5Value,
			want:   5,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.Length(); got != tt.want {
				t.Errorf("queue.Length() = %v, want %v", got, tt.want)
			}
		})
	}
}
