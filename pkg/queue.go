package pkg

type Queue interface {
	Enqueue(interface{})
	Dequeue() interface{}
}
