package test

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/require"
	q "queue/queue"
	"testing"
)

func TestQueue(t *testing.T) {

	Convey("TestQueue", t, func() {
		queue := q.New[int]()
		queue.Enqueue(1)
		queue.Enqueue(2)

		Convey("TestQueueEnqueue", func() {

			elementHead, _ := queue.Peek()
			So(elementHead, ShouldEqual, 1)
			queue.Enqueue(3)
			So(elementHead, ShouldEqual, 1)
		})

		Convey("TestQueueDequeue", func() {

			elementHead, _ := queue.Peek()
			So(elementHead, ShouldEqual, 1)
			deletedValue, err := queue.Dequeue()
			if err != nil {
				t.Error(err)
			}
			So(deletedValue, ShouldEqual, 1)
			So(err, ShouldBeNil)
			elementHead, _ = queue.Peek()
			So(elementHead, ShouldEqual, 2)
		})

		Convey("TestQueuePeek", func() {
			value, err := queue.Peek()
			if err != nil {
				t.Error(err)
			}
			So(value, ShouldEqual, 1)
		})
	})
}

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
