package pkg

import "testing"

func TestSliceQueue(t *testing.T) {
	q := &SliceQueue{}

	q.Enqueue(1)

	if item := q.Dequeue(); item != 1 {
		t.Error("Queue-Deque of one integer fails")
	}

	if item := q.Dequeue(); item != nil {
		t.Error("Queue-Deque of one integer fails when last must be nil")
	}
}

func BenchmarkSliceQueue(b *testing.B) {
	q := &SliceQueue{}

	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}

	for i := 0; i < b.N; i++ {
		q.Dequeue()
	}
}
