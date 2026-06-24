package main

type Queue struct {
	Head *Node
}

type Node struct {
	Val  string
	Next *Node
}

func (q *Queue) Enqueue(name string) {
	node := Node{Val: name}

	if q.Head == nil {
		q.Head = &node
	} else {
		q.Head.Next = &node
	}
}

func (q *Queue) Dequeue() string {
	if q.Head == nil {
		return ""
	}

	node := q.Head
	q.Head = q.Head.Next

	return node.Val
}
