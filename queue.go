package main

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

func (n *Queue) Add(data int) {
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

func (n *Queue) Pop() {
	if n.Start == nil {
		return
	}

	n.Start = n.Start.Next
	n.Size--

	if n.Start == nil {
		n.Tail = nil
	}
}
