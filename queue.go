package main

import "errors"

type Node struct {
	Data int
	Next *Node
}

type Queue struct {
	Size  int
	Start *Node
	Tail  *Node
}

func NewQueue() *Queue {
	return &Queue{
		Size:  0,
		Start: nil,
		Tail:  nil,
	}
}

func (n *Queue) Enqueue(data int) {
	temp := &Node{
		Data: data,
		Next: nil,
	}

	if n.Size == 0 {
		n.Start = temp
		n.Tail = temp
	} else {
		n.Tail.Next = temp
		n.Tail = temp
	}

	n.Size++
}

func (n *Queue) Dequeue() (int, error) {
	if n.Start == nil {
		panic(errors.New("Queue is empty"))
	}

	dequeuedNode := n.Start
	n.Start = n.Start.Next
	n.Size--

	if n.Start == nil {
		n.Tail = nil
	}

	return dequeuedNode.Data, nil
}
