package taskS

import "time"

type Task struct {
	name        string
	text        string
	state       bool
	CreatedTime time.Time
	FinishTime  *time.Time
}

func (t *Task) Completed() {
	t.state = true
	now := time.Now()
	t.FinishTime = &now
}

func NewTask(name, text string) Task {
	return Task{
		name:        name,
		text:        text,
		state:       false,
		CreatedTime: time.Now(),
		FinishTime:  nil,
	}

}
