package usecase

import (
	"errors"
	"task_manager/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskUsecase struct {
	TaskRepo domain.TaskRepositoryInterface
}

func NewTaskUsecase(repo domain.TaskRepositoryInterface) *TaskUsecase {
	return &TaskUsecase{
		TaskRepo: repo,
	}
}

func (u *TaskUsecase) GetAllTasks() ([]domain.Task, error) {
	tasks, err := u.TaskRepo.GetAllTasks()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (u *TaskUsecase) GetTaskByID(id primitive.ObjectID) (*domain.Task, error) {
	task, err := u.TaskRepo.GetTaskByID(id)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (u *TaskUsecase) CreateTask(task *domain.Task, claims *domain.Claims) (*domain.Task, error) {
	if claims == nil {
		return nil, errors.New("invalid user claims")
	}

	task.UserId = claims.ID.Hex()

	// Generate a new ID for the task
	task.ID = primitive.NewObjectID()

	err := u.TaskRepo.AddTask(task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (u *TaskUsecase) UpdateTask(id primitive.ObjectID, taskData bson.M)(*domain.Task,error) {
	task, err := u.TaskRepo.UpdateTask(id, taskData)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (u *TaskUsecase) DeleteTask(id primitive.ObjectID) error {
	return u.TaskRepo.DeleteTask(id)
}
