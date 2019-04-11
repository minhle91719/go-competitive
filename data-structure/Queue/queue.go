package queue

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

type IQueue interface {
	EnqueueWithSlice([]interface{})
	Enqueue(value interface{})
	Dequeue() (value interface{}, err error)

	ToSliceWithoutDequeue() []interface{}
	ToSliceWithDequeue() []interface{}

	Length() uint32
	Capacity() uint32
}

var (
	QueueEmpty error = errors.New("queue empty")

	TIMEOUT_ENQUEUE_AVAILABLE = 10 * time.Second
)

type queue struct {
	capacity    uint32
	lengthQueue uint32

	front *node
	rear  *node

	mu sync.Mutex
}
type node struct {
	value interface{}

	front *node
	rear  *node
}

func NewQueue(cap uint32) IQueue {
	return &queue{
		capacity:    cap,
		lengthQueue: 0,
		mu:          sync.Mutex{},
	}
}
func (q *queue) Length() uint32 {
	return atomic.LoadUint32(&q.lengthQueue)
}
func (q *queue) Capacity() uint32 {
	return atomic.LoadUint32(&q.capacity)
}

func (q *queue) Enqueue(value interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.enqueue(value)
}
func (q *queue) EnqueueWithSlice(args []interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()
	for _, v := range args {
		q.enqueue(v)
	}
}

func (q *queue) Dequeue() (interface{}, error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.isEmpty() {
		return nil, QueueEmpty
	}
	return q.dequeue(), nil
}
func (q *queue) ToSliceWithDequeue() []interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()

	var slice []interface{}
	for q.isEmpty() == false {
		slice = append(slice, q.dequeue())
	}
	return slice
}
func (q *queue) ToSliceWithoutDequeue() []interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.isEmpty() {
		return nil
	}
	var slice []interface{}

	tmpNode := q.front
	for tmpNode != nil {
		slice = append(slice, tmpNode.value)
		tmpNode = tmpNode.rear
	}
	return slice
}

func (q *queue) enqueue(value interface{}) {
	q.checkEnqueueAvailable()
	defer atomic.AddUint32(&q.lengthQueue, 1)
	nV := &node{
		value: value,
	}
	if q.isEmpty() {
		q.rear, q.front = nV, nV
		return
	}
	q.rear.rear, nV.front, q.rear = nV, q.rear, nV
}
func (q *queue) dequeue() interface{} {
	defer atomic.AddUint32(&q.lengthQueue, ^uint32(0))
	nv := q.front
	if q.front != nil {
		q.front = q.front.rear
	} else {
		q.front, q.rear = nil, nil
	}

	return nv.value
}
func (q *queue) isEmpty() bool {
	return atomic.LoadUint32(&q.lengthQueue) == 0
}

func (q *queue) checkEnqueueAvailable() {
	tick := time.NewTicker(TIMEOUT_ENQUEUE_AVAILABLE)
	for {
		select {
		case <-tick.C:
			panic("Enqueue available timeout.")
		default:
			if atomic.LoadUint32(&q.lengthQueue) < q.capacity {
				tick.Stop()
				return
			}
		}
	}
}
