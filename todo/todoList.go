package todo

import (
	"GoLess2/taskS"
	"fmt"
	"github.com/k0kubun/pp"
)

type List struct {
	list map[string]*taskS.Task
}

func NewList() List {
	return List{
		list: make(map[string]*taskS.Task),
	}
}

func (l *List) ShowTasks() {
	if len(l.list) == 0 {
		fmt.Println("No tasks")
		return
	}
	t := map[string]*taskS.Task{}
	t = l.list
	for k, v := range t {
		pp.Println(k, *v)
	}
}

func (l *List) AddTask(name, text string) {
	for k, _ := range l.list {
		if k == name {
			fmt.Println("Task already exists")
			return
		}
	}
	add := taskS.NewTask(name, text)
	l.list[name] = &add
}

func (l *List) DeleteTask(name string) {
	count := 0
	for k, _ := range l.list {
		if k == name {
			delete(l.list, name)
			count++
			fmt.Println("Task deleted")
		}
	}
	if count == 0 {
		fmt.Println("Task not found")
	}
}

func (l *List) CompletedTask(name string) {
	count := 0
	for k, _ := range l.list {
		if k == name {
			l.list[name].Completed()
			count++
			fmt.Println("Task completed")
		}
	}
	if count == 0 {
		fmt.Println("Task not found")
	}
}
