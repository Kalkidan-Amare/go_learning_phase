package usecase_test

import (
	"errors"
	"testing"
	"task_manager/domain"
	"task_manager/mocks"
	"task_manager/usecase"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskUsecaseTestSuite struct {
	suite.Suite
	mockRepo    *mocks.TaskRepositoryInterface
	taskUsecase *usecase.TaskUsecase
}

func (suite *TaskUsecaseTestSuite) SetupTest() {
	suite.mockRepo = new(mocks.TaskRepositoryInterface)
	suite.taskUsecase = usecase.NewTaskUsecase(suite.mockRepo)
}

func (suite *TaskUsecaseTestSuite) TestGetAllTasks() {
	// Success case
	suite.Run("GetAllTasks_Success", func() {
		tasks := []domain.Task{
			{ID: primitive.NewObjectID(), Title: "Task 1"},
			{ID: primitive.NewObjectID(), Title: "Task 2"},
		}

		suite.mockRepo.On("GetAllTasks").Return(tasks, nil).Once()

		result, err := suite.taskUsecase.GetAllTasks()
		suite.Nil(err)
		suite.Equal(tasks, result)
	})

	// Failure case
	suite.Run("GetAllTasks_Failure", func() {
		suite.mockRepo.On("GetAllTasks").Return(nil, errors.New("error")).Once()

		result, err := suite.taskUsecase.GetAllTasks()
		suite.NotNil(err)
		suite.Nil(result)
	})
}

func (suite *TaskUsecaseTestSuite) TestGetTaskByID() {
	// Success case
	suite.Run("GetTaskByID_Success", func() {
		taskID := primitive.NewObjectID()
		task := &domain.Task{ID: taskID, Title: "Task"}

		suite.mockRepo.On("GetTaskByID", taskID).Return(task, nil).Once()

		result, err := suite.taskUsecase.GetTaskByID(taskID)
		suite.Nil(err)
		suite.Equal(task, result)
	})

	// Failure case
	suite.Run("GetTaskByID_Failure", func() {
		taskID := primitive.NewObjectID()

		suite.mockRepo.On("GetTaskByID", taskID).Return(nil, errors.New("not found")).Once()

		result, err := suite.taskUsecase.GetTaskByID(taskID)
		suite.NotNil(err)
		suite.Nil(result)
	})
}

func (suite *TaskUsecaseTestSuite) TestCreateTask() {
	// Success case
	suite.Run("CreateTask_Success", func() {
		task := &domain.Task{Title: "New Task"}
		claims := &domain.Claims{ID: primitive.NewObjectID()}

		suite.mockRepo.On("AddTask", task).Return(primitive.NewObjectID(), nil).Once()

		result, err := suite.taskUsecase.CreateTask(task, claims)
		suite.Nil(err)
		suite.NotNil(result)
	})

	// Invalid claims case
	suite.Run("CreateTask_InvalidClaims", func() {
		task := &domain.Task{Title: "New Task"}

		result, err := suite.taskUsecase.CreateTask(task, nil)
		suite.NotNil(err)
		suite.Nil(result)
	})

	// Failure case
	suite.Run("CreateTask_Failure", func() {
		task := &domain.Task{Title: "New Task"}
		claims := &domain.Claims{ID: primitive.NewObjectID()}

		suite.mockRepo.On("AddTask", task).Return(nil, errors.New("failed to create task")).Once()

		result, err := suite.taskUsecase.CreateTask(task, claims)
		suite.NotNil(err)
		suite.Nil(result)
	})
}

func (suite *TaskUsecaseTestSuite) TestUpdateTask() {
	// Success case
	suite.Run("UpdateTask_Success", func() {
		taskID := primitive.NewObjectID()
		task := &domain.Task{Title: "Updated Task"}

		suite.mockRepo.On("UpdateTask", taskID, task).Return(task, nil).Once()

		result, err := suite.taskUsecase.UpdateTask(taskID, task)
		suite.Nil(err)
		suite.Equal(task, result)
	})

	// Failure case
	suite.Run("UpdateTask_Failure", func() {
		taskID := primitive.NewObjectID()
		task := &domain.Task{Title: "Updated Task"}

		suite.mockRepo.On("UpdateTask", taskID, task).Return(nil, errors.New("failed to update task")).Once()

		result, err := suite.taskUsecase.UpdateTask(taskID, task)
		suite.NotNil(err)
		suite.Nil(result)
	})
}

func (suite *TaskUsecaseTestSuite) TestDeleteTask() {
	// Success case
	suite.Run("DeleteTask_Success", func() {
		taskID := primitive.NewObjectID()

		suite.mockRepo.On("DeleteTask", taskID).Return(nil).Once()

		err := suite.taskUsecase.DeleteTask(taskID)
		suite.Nil(err)
	})

	// Failure case
	suite.Run("DeleteTask_Failure", func() {
		taskID := primitive.NewObjectID()

		suite.mockRepo.On("DeleteTask", taskID).Return(errors.New("failed to delete task")).Once()

		err := suite.taskUsecase.DeleteTask(taskID)
		suite.NotNil(err)
	})
}

func TestTaskUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(TaskUsecaseTestSuite))
}
