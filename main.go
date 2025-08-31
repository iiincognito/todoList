package main

import "GoLess2/todo"

func main() {
	ev := NewEventStore()
	todoList := todo.NewList()
	scanner := NewScanner(&todoList)
	scanner.Start(ev)
}
