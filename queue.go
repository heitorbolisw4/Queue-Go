package main

type Queue struct {
	Head *Node
}

type Node struct {
	Val  string
	Next *Node
}
