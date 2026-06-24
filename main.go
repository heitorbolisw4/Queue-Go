package main

import "fmt"

func main() {
	queue := Queue{}

	queue.Enqueue("Learning Golang")
	queue.Enqueue("Studying Data Structure")
	queue.Enqueue("Playing Counter-Strike")

	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
}
