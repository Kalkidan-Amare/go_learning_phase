package main

import (
    "context"
    "fmt"
    "log"
	"os"

	"github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI is not set in .env file")
	}

	//set client options
	clientOptions := options.Client().ApplyURI(mongoURI)

	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// check the connection
	err = client.Ping(context.TODO(),nil)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("connected to mongodb")

	collection := client.Database("test").Collection("trainers")

	ash := Trainer{"Ash", 10, "Pallet Town"}
	misty := Trainer{"Misty", 10, "Pallet Town"}
	brock := Trainer{"Brock", 15, "broklick"}


	//INSERT a single
	insert, err := collection.InsertOne(context.TODO(),ash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insert.InsertedID)


	//Insert a new document MANY
	trainers := []interface{}{misty,brock}
	insertMany, err := collection.InsertMany(context.TODO(),trainers)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted multiple documents: ", insertMany.InsertedIDs)


	// Updating data
	filter := bson.D{{"name","Ash"}}

	update := bson.D{
		{"$inc", bson.D{
			{"age", 1},
		}},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	
	//GET ONE
	var result Trainer

	err = collection.FindOne(context.TODO(), bson.D{{"name","Ash"}}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single documet: %+v\n", result)


	//GET ALL
	findOptions := options.Find()
	findOptions.SetLimit(2)

	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil{
		log.Fatal(err)
	}

	var results[]*Trainer

	for cur.Next(context.TODO()){
		var ele Trainer
		if err = cur.Decode(&ele); err != nil {
			log.Fatal(err)
		}

		results = append(results,&ele)	
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)

	
	//DELETE
	deleteResult, err := collection.DeleteOne(context.TODO(), bson.D{{"name","Brock"}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)

	//disconnecting the database
	err = client.Disconnect(context.TODO())
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed")
}