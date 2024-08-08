package data

import (
    "context"
    "errors"
    "task_manager/models"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection

func SetUserCollection(client *mongo.Client){
	userCollection = client.Database("TaskManager").Collection("Users")
}

func RegisterUser(user models.User) (models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}
	user.Password = string(hashedPassword)

	_,err = userCollection.InsertOne(context.Background(),user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func AuthenticateUser(newUser models.User) (models.User, error) {
	var user models.User
	err := userCollection.FindOne(context.Background(),bson.M{"username": newUser.Username}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(newUser.Password))
	if err != nil{
		return models.User{}, errors.New("invalide credentials")
	}

	return user,nil
}