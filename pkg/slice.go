package pkg

type SliceQueue struct {
	queue []interface{}
}

func (q *SliceQueue) Enqueue(item interface{}) {
	q.queue = append(q.queue, item)
}

func (q *SliceQueue) Dequeue() interface{} {
	if q.queue == nil {
		return nil
	}

	item := q.queue[0]

	if len(q.queue) > 1 {
		q.queue = q.queue[1:]
	} else {
		q.queue = nil
	}

	return item
}
