package config

import(
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init()(*mongo.Client,error){
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI is not set in .env file")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(context.TODO(),clientOptions)
	if err != nil {
		return &mongo.Client{}, nil
	}

	if err = client.Ping(context.TODO(),nil); err != nil {
		return &mongo.Client{}, nil
	}

	return client, nil
}