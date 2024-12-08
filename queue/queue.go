package queue

import "fmt"

type Node[A any] struct {
	Value A
	Next  *Node[A]
}

type Queue[A any] struct {
	Head *Node[A]
	Tail *Node[A]
}

func New[A any]() *Queue[A] {
	return &Queue[A]{}
}

func (q *Queue[A]) Peek() (A, error) {
	if q.Head == nil {
		return q.Head.Value,
			fmt.Errorf("queue is empty")
	}
	return q.Head.Value, nil
}

func (q *Queue[A]) Enqueue(value A) {
	node := &Node[A]{Value: value}
	if q.Head == nil {
		q.Head = node
	}
	if q.Tail != nil {
		q.Tail.Next = node
	}
	q.Tail = node
}

func (q *Queue[A]) Dequeue() (A, error) {
	if q.Head == nil {
		return q.Head.Value,
			fmt.Errorf("queue is empty")
	}
	value := q.Head.Value
	q.Head = q.Head.Next
	return value, nil
}
