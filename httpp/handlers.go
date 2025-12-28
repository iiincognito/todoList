package httpp

import (
	"GoLess2/taskS"
	"GoLess2/todo"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type HTTPHandlers struct {
	todoList *todo.List
}

func NewHTTPHandlers(todoList *todo.List) *HTTPHandlers {
	return &HTTPHandlers{
		todoList: todoList,
	}
}

/*
pattern: /tasks
method: POST
info: JSON

succeed:

	-status code: 201 Created
	-response body: JSON

failed:

	-status code: 400,409,500
	-response body: JSON with error + time error
*/
func (h *HTTPHandlers) HandleCreatedTask(w http.ResponseWriter, r *http.Request) {
	var taskDTO DTOTask
	if err := json.NewDecoder(r.Body).Decode(&taskDTO); err != nil {
		errDto := DTOError{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errDto.ToString(), http.StatusBadRequest)
		return
	}
	if err := taskDTO.ValidateForCreate(); err != nil {
		errDto := DTOError{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errDto.ToString(), http.StatusBadRequest)
	}
	todoTask := taskS.NewTask(taskDTO.Title, taskDTO.Description)
	if err := h.todoList.AddTask(todoTask); err != nil {
		errDto := DTOError{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errDto.ToString(), http.StatusConflict)
		return
	}
	b, err := json.MarshalIndent(todoTask, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(b)
	if err != nil {
		fmt.Println(err)
		return
	}

}

/*
pattern: /tasks{title}
method: GET
info: pattern

succeed:

	-status code: 200 ok
	-response body: JSON

failed:

	-status code: 400,404,500
	-response body: JSON with error + time error
*/
func (h *HTTPHandlers) HandleGetTask(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]
	task, err := h.todoList.GetTask(title)
	if err != nil {
		errDto := DTOError{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errDto.ToString(), http.StatusNotFound)
		return
	}
	b, err := json.MarshalIndent(task, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(b)
	if err != nil {
		fmt.Println(err)
		return
	}
}

/*
pattern: /tasks
method: GET
info: -

succeed:

	-status code: 200 ok
	-response body: JSON tasks

failed:

	-status code: 400,500
	-response body: JSON with error + time error
*/
func (h *HTTPHandlers) HandleGetAllTask(w http.ResponseWriter, r *http.Request) {
	tasks := h.todoList.ListTasks()
	b, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(b)
	if err != nil {
		fmt.Println(err)
		return
	}
}

/*
pattern: /tasks
method: GET
info: query params

succeed:

	-status code: 200 ok
	-response body: JSON tasks

failed:

	-status code: 400,500
	-response body: JSON with error + time error
*/
func (h *HTTPHandlers) HandleGetAllUncompletedTask(w http.ResponseWriter, r *http.Request) {
	unComplTasks := h.todoList.ListNotCompletedTasks()
	b, err := json.MarshalIndent(unComplTasks, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(b)
	if err != nil {
		fmt.Println(err)
		return
	}
}

/*
pattern: /tasks/{title}
method: PATCH
info: pattern + JSON

succeed:

	-status code: 200 ok
	-response body: JSON tasks

failed:

	-status code: 400,500
	-response body: JSON with error + time error
*/
func (h *HTTPHandlers) HandleCompleteTask(w http.ResponseWriter, r *http.Request) {
	var DTOCom DTOComplete
	err := json.NewDecoder(r.Body).Decode(&DTOCom)
	if err != nil {
		errDto := DTOError{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errDto.ToString(), http.StatusBadRequest)
		return
	}
	title := mux.Vars(r)["title"]
	if DTOCom.Completed {
		err = h.todoList.CompleteTask(title)
		if err != nil {
			errDto := DTOError{
				Message: err.Error(),
				Time:    time.Now(),
			}
			http.Error(w, errDto.ToString(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("task completed successfully"))
	} else {

	}
}

/*
pattern: /tasks/{title}
method: DELETE
info: pattern

succeed:

	-status code: 204 no content
	-response body: -

failed:

	-status code: 400,500
	-response body: JSON with error + time error
*/
func (h *HTTPHandlers) HandleDeleteTask(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]
	err := h.todoList.DeleteTask(title)
	if err != nil {
		errDto := DTOError{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errDto.ToString(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("task deleted successfully"))
	if err != nil {
		fmt.Println(err)
	}

}
