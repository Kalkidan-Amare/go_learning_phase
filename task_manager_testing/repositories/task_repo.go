package repositories

import (
	"context"
	// "fmt"
	"task_manager/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository struct {
	collection domain.Collection
}

func NewTaskRepository(db domain.Collection) *TaskRepository {
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

func (r *TaskRepository) UpdateTask(id primitive.ObjectID, taskData *domain.Task) (*domain.Task, error) {
    filter := bson.M{"_id": id}
    update := bson.M{"$set": taskData}
    
    result := r.collection.FindOneAndUpdate(context.Background(), filter, update)
    if result.Err() != nil {
        return nil, result.Err()
    }

    var decoded domain.Task
    if err := result.Decode(&decoded); err != nil {
        return nil, err
    }

	taskData.ID = decoded.ID
	taskData.UserId = decoded.UserId
	// fmt.Println(decoded.UserId)

	// updatedTask := domain.Task{
	// 	ID:          decoded.ID,
	// 	Title:       taskData["title"].(string),
	// 	Description: taskData["description"].(string),
	// 	DueDate:     taskData["due_date"].(string),
	// 	Status:      taskData["status"].(string),
	// 	UserId:      decoded.UserId,
	// }

    return taskData, nil
}


func (r *TaskRepository) DeleteTask(id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}
