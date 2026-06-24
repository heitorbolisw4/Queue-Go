package main

import "fmt"

func main() {
	queue := Queue{}

	queue.Enqueue("Heitor")
	queue.Enqueue("Golang")
	queue.Enqueue("Counter-Strike")

	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
}
