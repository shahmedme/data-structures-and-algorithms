package main

import "fmt"

var items [10]int
var front int = -1
var rear int = -1

func main() {
	deQueue()
	enQueue(1)
	enQueue(2)
	enQueue(3)
	enQueue(4)
	enQueue(5)
	enQueue(6)
	display()
	deQueue()
	deQueue()
	display()
}

func enQueue(item int) {
	if rear == 10-1 {
		fmt.Println("Queue is full")
	} else {
		if front == -1 {
			front = 0
		}
		rear++
		items[rear] = item
		fmt.Println("Item is Enqueued")
	}
}

func deQueue() {
	if front == -1 {
		fmt.Println("Queue is empty")
	} else {
		front++
		fmt.Println("Item is Dequeued")
		if front > rear {
			front = -1
			rear = -1
		}
	}
}

func display() {
	if rear == -1 {
		fmt.Println("Queue is empty")
	} else {
		for i := front; i <=rear; i++ {
			fmt.Print(items[i], " ")
		}
		fmt.Printf("\n")
	}
}