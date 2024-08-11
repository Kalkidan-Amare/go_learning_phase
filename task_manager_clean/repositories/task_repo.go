package repositories

import (
	"context"
	"errors"
	"task_manager/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository struct {
	collection *mongo.Collection
}

func NewTaskRepository(db *mongo.Collection) *TaskRepository {
	return &TaskRepository{
		collection: db,
	}
}

func (r *TaskRepository) GetAllTasks() ([]domain.Task, error) {
	var all []domain.Task
	cursor, err := r.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	err = cursor.All(context.TODO(),&all)

	return all, err
}

func (r *TaskRepository) GetTaskByID(id primitive.ObjectID) (*domain.Task, error) {
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	var task domain.Task
	err := r.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *TaskRepository) AddTask(task *domain.Task) (interface{},error) {
	inserted, err := r.collection.InsertOne(context.TODO(), task)
	return inserted.InsertedID,err
}

func (r *TaskRepository) UpdateTask(id primitive.ObjectID, updatedTask bson.M) (interface{},error) {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": updatedTask}

	task, err := r.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	if task.MatchedCount == 0 {
		return nil, errors.New("no task found with the given ID")
	}

	return task, nil
}

func (r *TaskRepository) DeleteTask(id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}
