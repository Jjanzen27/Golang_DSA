package main

import (
	"fmt"
)

type Node struct {
	Value interface{} // All types satisfy the empty interface, so we can store anything here.
	Next  *Node
}
type Queue struct {
	Front *Node
	Rear  *Node
	Size  int
}

func (q *Queue) Enqueue(value interface{}) {
	newNode := Node{value, nil}
	if q.Front == nil && q.Rear == nil {
		q.Front = &newNode
		q.Rear = &newNode
	} else {
		q.Rear.Next = &newNode
		q.Rear = &newNode
	}
	q.Size += 1
}

func (q *Queue) Dequeue() interface{} {
	if q.Front == nil {
		return nil
	} else {
		temp := q.Front
		q.Front = q.Front.Next
		temp.Next = nil
		q.Size -= 1
		return temp.Value
	}
}

func (q *Queue) getSize() int {
	return q.Size
}

func main() {
	var q Queue
	q.Enqueue("hello")
	fmt.Println(q.getSize())
	q.Enqueue("world")
	fmt.Println(q.getSize())
	q.Dequeue()
	fmt.Println(q.getSize())
}
