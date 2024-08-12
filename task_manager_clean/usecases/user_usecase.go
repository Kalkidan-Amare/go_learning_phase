package usecase

import (
	"errors"
	// "fmt"
	"task_manager/domain"
	"task_manager/infrastructure"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecase struct {
	UserRepo domain.UserRepositoryInterface
}

func NewUserUsecase(repo domain.UserRepositoryInterface) *UserUsecase {
	return &UserUsecase{
		UserRepo: repo,
	}
}

func (u *UserUsecase) CreateUser(user *domain.User) (interface{}, error) {
	existingUser, _ := u.UserRepo.GetUserByUsername(user.Username)
	if existingUser != nil {
		return nil, errors.New("username already exists")
	}

	hashedPassword, err := infrastructure.HashPassword(user.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}
	user.Password = hashedPassword

	insertedID,err := u.UserRepo.AddUser(user)
	if err != nil {
		return nil, err
	}
	// fmt.Println(insertedID)
	// Cast insertedID to ObjectID and set it to the user's ID
	if objectId, ok := insertedID.(primitive.ObjectID); ok {
		user.ID = objectId
	}

	return user, nil
}

func (u *UserUsecase) LoginUser(authUser *domain.AuthUser) (string, error) {
	user, err := u.UserRepo.GetUserByUsername(authUser.Username)
	if err != nil {
		return "", err
	}

	
	err = infrastructure.ComparePassword(user.Password, authUser.Password)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := infrastructure.GenerateJWT(user)
	if err != nil {
		return "", err
	}

	return token, nil
}


func (u *UserUsecase) GetUserByID(objectID primitive.ObjectID) (*domain.User, error) {
	user, err := u.UserRepo.GetUserByID(objectID)
	if err != nil {
		return nil, err
	}
	
	return user, nil
}
