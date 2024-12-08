package queue

type QueueInterface[A any] interface {
	Enqueue(A)
	Dequeue() (A, error)
	Peek() (A, error)
}
