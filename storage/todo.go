package storage

import (
	"context"
	"gemini/helper"
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// Todo ..
type Todo struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Reminder    time.Time          `bson:"reminder"`
}

// TodoStorage ..
type TodoStorage interface {
	Store(data *Todo) (interface{}, error)
	// ByID(id int) (*Todo, error)
	List() ([]Todo, error)
}

type todoImpl struct {
	coll *mongo.Collection
}

// NewTodoStorage ..
func NewTodoStorage() TodoStorage {
	return &todoImpl{coll: db.Collection("todo")}
}

func (s *todoImpl) Store(data *Todo) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	inserted, err := s.coll.InsertOne(ctx, data)
	if err != nil {
		helper.Logging("Todo", "Store", err.Error())
		return nil, err
	}
	return inserted.InsertedID, nil
}

func (s *todoImpl) List() ([]Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := s.coll.Find(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)

	todos := []Todo{}
	for cur.Next(ctx) {
		var todo Todo
		if err := cur.Decode(&todo); err != nil {
			helper.Logging("Todo", "List", err.Error())
			return nil, err
		}

		todos = append(todos, todo)
	}
	if err := cur.Err(); err != nil {
		helper.Logging("Todo", "List", err.Error())
		return nil, err
	}

	return todos, nil
}
