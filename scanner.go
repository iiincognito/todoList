package main

import (
	"GoLess2/todo"
	"bufio"
	"github.com/k0kubun/pp"

	"fmt"
	"os"
	"strings"
)

type Scanner struct {
	todoList *todo.List
}

func NewScanner(todoList *todo.List) *Scanner {
	return &Scanner{
		todoList: todoList,
	}
}
func (s *Scanner) Start(e *EventStore) {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Enter command")
		ok := scanner.Scan()
		if !ok {
			fmt.Println("Scanner error")
			return
		}
		input := scanner.Text()
		err, res := s.process(input)
		e.AddEvent(NewEvent(res, input, err))
		if err != nil {
			fmt.Println(err)
			return
		}
		if res == "exit" {
			return
		}
		if res == "list" {
			fmt.Println("list of tasks")
			continue
		}
		if res == "del" {
			fmt.Println("delete task")
			continue
		}
		if res == "add" {
			fmt.Println("add task")
			continue
		}
		if res == "complete" {
			fmt.Println("complete task")
			continue
		}
		if res == "help" {
			fmt.Println("add\n" +
				"del\n" +
				"list\n" +
				"complete\n" +
				"exit")
			continue
		}
		if res == "events" {
			es := e.GetEvents()
			fmt.Println("events")
			for _, v := range es {
				pp.Println(v)
			}
			continue
		}
	}
}
func (s *Scanner) process(input string) (error, string) {
	fields := strings.Fields(input)

	if len(fields) == 0 {
		err := fmt.Errorf("Invalid input")
		return err, ""
	}

	cmd := fields[0]

	if cmd == "exit" {
		return nil, "exit"
	}

	if cmd == "help" {
		return nil, "help"
	}

	if cmd == "events" {
		return nil, "events"
	}

	if cmd == "add" {
		return s.AddTask(fields), "add"
	}

	if cmd == "del" {
		return s.DeleteTask(fields), "del"
	}

	if cmd == "list" {
		s.todoList.ShowTasks()
		return nil, "list"
	}
	if cmd == "complete" {
		return s.CompleteTask(fields), "complete"
	}

	return nil, ""
}

func (s *Scanner) CompleteTask(fields []string) error {
	if len(fields) < 2 {
		err := fmt.Errorf("Invalid input, need 2 arguments")
		return err
	}
	task := fields[1]
	s.todoList.CompletedTask(task)
	return nil
}

func (s *Scanner) AddTask(fields []string) error {
	if len(fields) < 3 {
		err := fmt.Errorf("Invalid input, need 3 arguments")
		return err
	}
	task := fields[1]
	text := ""
	for i := 2; i < len(fields); i++ {
		text += fields[i]
		if i != len(fields)-1 {
			text += " "
		}
	}
	s.todoList.AddTask(task, text)
	return nil
}

func (s *Scanner) DeleteTask(fields []string) error {
	if len(fields) < 2 {
		err := fmt.Errorf("Invalid input, need 2 arguments")
		return err
	}
	task := fields[1]
	s.todoList.DeleteTask(task)
	return nil
}
