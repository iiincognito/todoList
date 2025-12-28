package httpp

import (
	"encoding/json"
	"errors"
	"time"
)

type DTOTask struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type DTOError struct {
	Message string
	Time    time.Time
}

type DTOComplete struct {
	Completed bool
}

func (t DTOTask) ValidateForCreate() error {
	if t.Title == "" {
		return errors.New("title is empty")
	}
	if t.Description == "" {
		return errors.New("text is empty")
	}
	return nil
}

func (e DTOError) ToString() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(b)
}
