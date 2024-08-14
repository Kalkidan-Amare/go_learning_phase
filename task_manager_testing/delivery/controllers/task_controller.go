package controllers

import (
	"net/http"
	"task_manager/domain"
	// "task_manager/infrastructure"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskController struct {
	TaskUsecase domain.TaskUsecaseInterface
}

func NewTaskController(usecase domain.TaskUsecaseInterface) *TaskController {
	return &TaskController{
		TaskUsecase: usecase,
	}
}

func (tc *TaskController) GetTasks(c *gin.Context) {
	tasks, err := tc.TaskUsecase.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (tc *TaskController) GetTask(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	task, err := tc.TaskUsecase.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	var newTask domain.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// logged-in user claims
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	claims, ok := user.(*domain.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user claims"})
		return
	}

	// Create the task with the logged-in user ID
	createdTaskId, err := tc.TaskUsecase.CreateTask(&newTask, claims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdTaskId)
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var updatedTask domain.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, _ := c.Get("user")
	claims, _ := user.(*domain.Claims)

	task, err := tc.TaskUsecase.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}



	if task.UserId != claims.ID.Hex() && claims.Role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only update your tasks"})
		return
	}

	// update := bson.M{
	// 	"title":       updatedTask.Title,
	// 	"description": updatedTask.Description,
	// 	"due_date":    updatedTask.DueDate,
	// 	"status":      updatedTask.Status,
	// }

	result,err := tc.TaskUsecase.UpdateTask(id, &updatedTask); 
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	user, _ := c.Get("user")
	claims, _ := user.(*domain.Claims) // Assuming Claims struct is in domain package

	task, err := tc.TaskUsecase.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if task.UserId != claims.ID.Hex() && claims.Role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only delete your tasks"})
		return
	}

	if err := tc.TaskUsecase.DeleteTask(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
