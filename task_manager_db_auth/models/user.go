package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username string             `bson:"username" json:"username"`
    Password string             `bson:"password" json:"password"`
    Role     string             `bson:"role" json:"role"`
}