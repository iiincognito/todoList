package main

import (
	"GoLess2/httpp"
	"GoLess2/todo"
)

func main() {
	todoList := todo.NewList()
	httpHand := httpp.NewHTTPHandlers(todoList)
	server := httpp.NewHTTPServer(httpHand)
	err := server.StartServer()
	if err != nil {
		panic(err)
	}
}
