package taskS

import "time"

type Task struct {
	Name        string
	text        string
	Completed   bool
	CreatedTime time.Time
	FinishTime  *time.Time
}

func (t *Task) Complete() {
	t.Completed = true
	now := time.Now()
	t.FinishTime = &now
}

func (t *Task) Uncomplete() {
	t.Completed = false
	t.FinishTime = nil
}

func NewTask(name, text string) Task {
	return Task{
		Name:        name,
		text:        text,
		Completed:   false,
		CreatedTime: time.Now(),
		FinishTime:  nil,
	}

}
