package main

import (
	"fmt"
)

type queue struct {
	arr []int
}

func (q *queue) enqueue(x int) {
	(*q).arr = append((*q).arr, x)
}

func (q *queue) dequeue() int {
	if len((*q).arr) == 0 {
		return -1
	}
	k := (*q).arr[0]
	(*q).arr = (*q).arr[1:len((*q).arr)]

	return k
}

func main() {
	fmt.Println("Queue Program")
	q1 := queue{}
	q1.enqueue(10)
	q1.enqueue(20)
	q1.enqueue(30)
	fmt.Println(q1)
	_ = q1.dequeue()
	fmt.Println(q1, "first")
	_ = q1.dequeue()
	fmt.Println(q1, "second")
	_ = q1.dequeue()
	fmt.Println(q1)
	q1.enqueue(40)
	fmt.Println(q1)
}
