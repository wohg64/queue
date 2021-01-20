package pkg

import "testing"

func TestListQueue(t *testing.T) {
	q := New()

	q.Enqueue(1)

	if item := q.Dequeue(); item != 1 {
		t.Error("Queue-Deque of one integer fails")
	}

	if item := q.Dequeue(); item != nil {
		t.Error("Queue-Deque of one integer fails when last must be nil")
	}
}

func BenchmarkListQueue(b *testing.B) {
	q := New()

	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}

	for i := 0; i < b.N; i++ {
		q.Dequeue()
	}
}
