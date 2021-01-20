package pkg

import "container/list"

type ListQueue struct {
	queue *list.List
}

func New() *ListQueue {
	return &ListQueue{
		queue: list.New(),
	}
}

func (q *ListQueue) Enqueue(item interface{}) {
	q.queue.PushBack(item)
}

func (q *ListQueue) Dequeue() interface{} {
	item := q.queue.Front()
	if item == nil {
		return nil
	}

	q.queue.Remove(item)
	return item.Value
}
