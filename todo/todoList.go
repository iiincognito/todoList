package todo

import (
	. "GoLess2/taskS"
	"errors"
	"fmt"
	"sync"
)

type List struct {
	list map[string]*Task
	mtx  sync.RWMutex
}

func NewList() *List {
	return &List{
		list: make(map[string]*Task),
	}
}

func (l *List) ListTasks() map[string]*Task {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	if len(l.list) == 0 {
		fmt.Println("No tasks")
		return nil
	}
	t := map[string]*Task{}
	t = l.list
	return t
}

func (l *List) AddTask(task Task) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	if _, ok := l.list[task.Name]; ok {
		fmt.Println("Task already exists")
		return errors.New("Task already exists")
	}
	l.list[task.Name] = &task
	return nil
}

func (l *List) GetTask(name string) (Task, error) {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	res, ok := l.list[name]
	if !ok {
		return Task{}, errors.New("Task not found")
	}

	return *res, nil
}

func (l *List) ListNotCompletedTasks() map[string]Task {
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	notCompletedTasks := map[string]Task{}
	for k, v := range l.list {
		if v.Completed == false {
			notCompletedTasks[k] = *v
		}
	}
	return notCompletedTasks
}

func (l *List) DeleteTask(name string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	_, ok := l.list[name]
	if !ok {
		return errors.New("Task not found")
	}
	delete(l.list, name)
	fmt.Println("Task deleted")
	return nil
}

func (l *List) CompleteTask(name string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	count := 0
	for k, _ := range l.list {
		if k == name {
			l.list[name].Complete()
			count++
			fmt.Println("Task completed")
		}
	}
	if count == 0 {
		return errors.New("Task not found")
	}
	return nil
}

func (l *List) UncompleteTask(name string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	res, ok := l.list[name]
	if !ok {
		return errors.New("Task not found")
	}
	if res.Completed == true {
		l.list[name].Uncomplete()
	}
	return nil
}
