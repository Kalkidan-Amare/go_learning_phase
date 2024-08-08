package data

import (
	"context"
	"errors"
	"task_manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// var tasks = make(map[int]models.Task)
// var currID = 1
var taskCollection *mongo.Collection


func SetTaskCollection(client *mongo.Client){
	taskCollection = client.Database("TaskManager").Collection("Tasks")
}

func GetAllTasks() ([]models.Task,error) {
	var all []models.Task

	cur,err := taskCollection.Find(context.TODO(),bson.M{})
	if err != nil {
		return []models.Task{},err
	}

	err = cur.All(context.TODO(),&all)
	if err != nil {
		return []models.Task{}, err
	}

	return all,nil
}

func GetTaskByID(id primitive.ObjectID) (models.Task, error) {
	var task models.Task
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	err := taskCollection.FindOne(context.TODO(),filter).Decode(&task)
	if err != nil {
		return models.Task{},err
	}

	return task, nil
}

func CreateTask(task models.Task) (interface{}, error) {
	insert, err := taskCollection.InsertOne(context.TODO(),task)
	if err != nil {
		return models.Task{},err
	}

	return insert.InsertedID,nil
}

func UpdateTask(id primitive.ObjectID, updatedTask models.Task) (models.Task, error) {
	filter := bson.M{"_id": id}
	
	update := bson.M{"$set": updatedTask}

	task, err := taskCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return models.Task{}, err
	}
	if task.MatchedCount == 0 {
		return models.Task{}, errors.New("no task found with the given ID")
	}

	return updatedTask, nil
}

func DeleteTask(id primitive.ObjectID) error{
	_,err := taskCollection.DeleteOne(context.TODO(),bson.D{{Key: "_id", Value: id}})
	if err != nil{
		return err
	}

	return nil
}