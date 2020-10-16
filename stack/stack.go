package main

import "fmt"

type stack struct {
	items [10]int
	top int
}

func main() {
	var st *stack = new(stack)
	st.top = -1

	push(st, 1)
	push(st, 2)
	push(st, 3)
	push(st, 4)
	pop(st)
	printStack(st)
}

func printStack(st *stack) {
	for i := 0; i <= st.top; i++ {
		fmt.Println(st.items[i])
	}
} 

func push(st *stack, item int) {
	if isFull(st) {
		fmt.Println("Stack is full")
	} else {
		st.top++
		st.items[st.top] = item
	}
}

func pop(st *stack) {
	if isEmpty(st) {
		fmt.Println("Stack is empty")
	} else {
		st.top--
	}
}

func isFull(st *stack) bool {
	if st.top == 10 {
		return true
	}
	return false
}

func isEmpty(st *stack) bool {
	if st.top == 0 {
		return true
	}
	return false
}