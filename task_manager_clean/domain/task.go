package domain

import (
	// "context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct{
	ID primitive.ObjectID `json:"id" bson:"_id"`
	Title string `json:"title"`
	Description string `json:"description"`
	DueDate string `json:"due_date"`
	Status string `json:"status"`
	UserId string `json:"user_id"`
}

type TaskUsecaseInterface interface {
	GetAllTasks() ([]Task, error)
	GetTaskByID(id primitive.ObjectID) (*Task, error)
	CreateTask(task *Task, claims *Claims) (*Task, error) 
	UpdateTask(id primitive.ObjectID, taskData bson.M) error
	DeleteTask(id primitive.ObjectID) error
}


type TaskRepositoryInterface interface {
	GetAllTasks() ([]Task, error)
	GetTaskByID(id primitive.ObjectID) (*Task, error)
	AddTask(task *Task) error
	UpdateTask(id primitive.ObjectID, taskData bson.M) (*Task,error)
	DeleteTask(id primitive.ObjectID) error
}