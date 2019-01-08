package handler

import (
	"context"
	"encoding/json"
	"gemini/pb/v1/todo"
	"gemini/service"
	"log"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// TodoHandler ..
type TodoHandler interface {
	Store(ctx context.Context, in *todo.StoreRequest) (*todo.StoreResponse, error)
	List(ctx context.Context, filter *todo.ListRequest) (*todo.ListResponse, error)
}

type todoHandlerImpl struct {
	todoSC service.TodoService
}

// NewTodoHandler ..
func NewTodoHandler() TodoHandler {
	return &todoHandlerImpl{
		todoSC: service.NewTodoService(),
	}
}

// Store ..
func (h *todoHandlerImpl) Store(ctx context.Context, request *todo.StoreRequest) (*todo.StoreResponse, error) {
	log.Println("Creating todo")
	response := h.todoSC.Store(request)
	if response.Err != nil {
		return nil, response.Err
	}
	return &todo.StoreResponse{
		Api:     apiVersion,
		Id:      response.Data.(primitive.ObjectID).Hex(),
		Success: true,
	}, nil
}

// List ..
func (h *todoHandlerImpl) List(ctx context.Context, filter *todo.ListRequest) (*todo.ListResponse, error) {
	log.Println("List todos")
	response := h.todoSC.List()
	if response.Err != nil {
		return nil, response.Err
	}

	var data []*todo.ToDo
	if err := json.Unmarshal(response.Data.([]byte), &data); err != nil {
		return nil, err
	}
	return &todo.ListResponse{
		Api:   apiVersion,
		ToDos: data,
	}, nil
}
