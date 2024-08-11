package repositories

import (
	"context"
	"task_manager/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Collection) *UserRepository {
	return &UserRepository{
		collection: db,
	}
}

func (r *UserRepository) AddUser(user *domain.User) error {
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	_, err := r.collection.InsertOne(context.TODO(), user)
	return err
}

func (r *UserRepository) GetUserByUsername(username string)(*domain.User, error) {
	var user *domain.User
	err := r.collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func(r *UserRepository) GetUserByID(id primitive.ObjectID)(*domain.User, error) {
	var user *domain.User
	err := r.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user, err
}