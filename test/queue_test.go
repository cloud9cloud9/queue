package test

import (
	"github.com/stretchr/testify/require"
	q "queue/queue"
	"testing"
)

func TestQueuePeek(t *testing.T) {

	queue := q.New[int]()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	value, err := queue.Peek()
	if err != nil {
		t.Error(err)
	}
	require.Equal(t, 1, value)
}

func TestQueueDequeue(t *testing.T) {
	queue := q.New[int]()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	value, err := queue.Dequeue()
	if err != nil {
		t.Error(err)
	}

	element, _ := queue.Peek()
	require.Equal(t, 1, value)
	require.Equal(t, 2, element)
}

func TestQueueEnqueue(t *testing.T) {
	queue := q.New[int]()
	queue.Enqueue(1)
	queue.Enqueue(2)

	element, _ := queue.Peek()
	require.Equal(t, 1, element)
	queue.Enqueue(3)
	require.Equal(t, 1, element)
}
