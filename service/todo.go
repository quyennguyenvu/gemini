package service

import (
	"encoding/json"
	"gemini/pb/v1/todo"
	"gemini/storage"
	"net/http"
	"time"
)

// TodoService ..
type TodoService interface {
	Store(request *todo.StoreRequest) Response
	// ByID(strID string) Response
	List() Response
}

// todoHandlerImpl ..
type todoServiceImpl struct {
	todoDS storage.TodoStorage
}

// NewTodoService ..
func NewTodoService() TodoService {
	return &todoServiceImpl{
		todoDS: storage.NewTodoStorage(),
	}
}

// Store ..
func (sc *todoServiceImpl) Store(request *todo.StoreRequest) Response {
	reminder, _ := time.Parse("2006-01-02 15:04:05", request.Reminder)
	todo := storage.Todo{
		Title:       request.Title,
		Description: request.Description,
		Reminder:    reminder,
	}

	inserted, err := sc.todoDS.Store(&todo)
	if err != nil {
		return Response{
			Data: nil,
			Code: http.StatusInternalServerError,
			Err:  err,
		}
	}

	return Response{
		Data: inserted,
		Code: http.StatusOK,
		Err:  nil,
	}
}

// List ..
func (sc *todoServiceImpl) List() Response {
	todos, err := sc.todoDS.List()
	if err != nil || todos == nil {
		return Response{
			Data: nil,
			Code: http.StatusUnprocessableEntity,
			Err:  err,
		}
	}

	res, err := json.Marshal(todos)
	if err != nil {
		return Response{
			Data: nil,
			Code: http.StatusInternalServerError,
			Err:  err,
		}
	}

	return Response{
		Data: res,
		Code: http.StatusOK,
		Err:  nil,
	}
}
