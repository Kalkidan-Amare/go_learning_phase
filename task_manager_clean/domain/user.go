package domain

import (
	// "context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	ID             primitive.ObjectID `json:"id"`
	Username       string             `json:"username"`
	// Password       string             `json:"password"`
	Role           string             `json:"role"`
	jwt.StandardClaims
}

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string             `json:"username"`
	Password string             `json:"password"`
	Role     string             `json:"role"`
}

// type UserHidden struct {
// 	ID       primitive.ObjectID `json:"id" bson:"_id"`
// 	Username string             `json:"username"`
// 	Password string             `json:"-"`
// 	Role     string             `json:"role"`
// }
type AuthUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserUsecaseInterface interface {
	CreateUser(user *User) (interface{}, error)
	LoginUser(user *AuthUser) (string, error)
}

type UserRepositoryInterface interface {
	AddUser(user *User) (interface{},error)
	GetUserByUsername(username string) (*User, error)
	GetUserByID(objectID primitive.ObjectID) (*User, error)
}

// type UserRepository interface {
// 	AddUser(user *User) error
// 	GetUsers() ([]User, error)
// 	GetUserByID(objectID primitive.ObjectID) (*User, error)
// 	GetUserByUsername(username string) (*User, error)
// 	UpdateUser(objectID primitive.ObjectID, userData bson.M) error
// 	DeleteUser(objectID primitive.ObjectID) error
// }