package Queue

// first in first out

type Data struct {
	data int
	prev *Data
}

type Queue struct {
	front *Data
	rear  *Data
}

func NewQueue() *Queue {
	return &Queue{
		front: nil,
		rear:  nil,
	}
}

func (q *Queue) Dequeue() Data {
	data := q.front
	q.front = q.front.prev
	return *data
}

func (q *Queue) EnQueue(data Data) {
	if q.rear == nil {
		q.front = &data
		q.rear = &data
		return
	}
	q.rear.prev = &data
	q.rear = &data
}

func (q *Queue) RemoveAtPosition(index int) {
	i := 0
	tmp := q.rear
	for tmp != nil {
		if i+1 == index {
			tmp.prev = tmp.prev.prev
			return
		}
		i++
	}
}
