package main

import (
	"fmt"
	q "queue/queue"
)

func main() {

	queue := q.New[int]()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	fmt.Println(queue.Peek())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Peek())
}
