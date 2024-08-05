package data

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	// "net/http"
	"task_manager/models"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var tasks = make(map[int]models.Task)
var currID = 1
var collection *mongo.Collection

func init(){
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGODB_URI is not set in .env file")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(context.TODO(),clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	if err = client.Ping(context.TODO(),nil); err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Connected to Mongo")

	collection = client.Database("TaskManager").Collection("Tasks")
}


func GetAllTasks() ([]models.Task,error) {
	var all []models.Task

	cur,err := collection.Find(context.TODO(),bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	err = cur.All(context.TODO(),&all)
	if err != nil {
		return []models.Task{}, err
	}
	// for cur.Next(context.TODO()) {
	// 	var temp models.Task
	// 	if err = cur.Decode(&temp); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	all = append(all,temp)
	// }

	// if err = cur.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	return all,nil
}

func GetTaskByID(id primitive.ObjectID) (models.Task, error) {
	var task models.Task
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	err := collection.FindOne(context.TODO(),filter).Decode(&task)
	if err != nil {
		return models.Task{},err
	}

	return task, nil
}

func CreateTask(task models.Task) (interface{}, error) {
	insert, err := collection.InsertOne(context.TODO(),task)
	if err != nil {
		return models.Task{},err
	}

	return insert.InsertedID,nil
}

// func UpdateTask(id primitive.ObjectID, updatedTask models.Task) (models.Task, error){
// 	filter := bson.M{"_id": id}
	
// 	update := bson.D{
// 		{"$set", bson.D{
// 			{"title", updatedTask.Title},
// 			{"description", updatedTask.Description},
// 			{"due_date", updatedTask.DueDate},
// 			{"status", updatedTask.Status},
// 		}},
// 	}
	

// 	task, err := collection.UpdateOne(context.TODO(),update,filter)
// 	if err != nil {
// 		return models.Task{},err
// 	}
// 	if task.MatchedCount == 0 {
// 		return models.Task{}, errors.New("No task found with the given ID")
// 	}

// 	return updatedTask,nil
// }

func UpdateTask(id primitive.ObjectID, updatedTask models.Task) (models.Task, error) {
	filter := bson.M{"_id": id}
	
	update := bson.D{
		{"$set", bson.D{
			{"title", updatedTask.Title},
			{"description", updatedTask.Description},
			{"due_date", updatedTask.DueDate},
			{"status", updatedTask.Status},
		}},
	}

	// Ensure correct order and parameters for UpdateOne
	task, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return models.Task{}, err
	}
	if task.MatchedCount == 0 {
		return models.Task{}, errors.New("No task found with the given ID")
	}

	return updatedTask, nil
}

func DeleteTask(id primitive.ObjectID) error{
	// filter := bson.D("id",id)
	_,err := collection.DeleteOne(context.TODO(),bson.D{{Key: "_id", Value: id}})
	if err != nil{
		return err
	}

	return nil
}