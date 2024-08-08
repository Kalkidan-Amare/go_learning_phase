package controllers

import (
	"net/http"
	"task_manager/data"
	"task_manager/models"
	"task_manager/middleware"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func GetTasks(c *gin.Context){
	all,err := data.GetAllTasks()
	if err != nil{
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(200, all)
}

func GetTask(c *gin.Context){
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
        return
    }

	task,err := data.GetTaskByID(id)
	if err != nil{
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, task)
}

func CreatTask(c *gin.Context){
	var newtask models.Task
	if err := c.ShouldBindJSON(&newtask); err != nil{
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	
	user,_ := c.Get("user")
	claims,ok := user.(*middleware.Claims)
	if !ok {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
        return
    }

	newUserID := claims.UserID

	newtask.UserID = newUserID

	id,err := data.CreateTask(newtask)
	if err != nil{
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, id)
}

func UpdateTask(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
		return
	}

	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	// updatedTask.ID = id
	
	user,_ := c.Get("user")
	claims := user.(*middleware.Claims)
	
	task, err := data.GetTaskByID(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if task.UserID != claims.UserID && claims.Role != "admin" {
        c.JSON(http.StatusForbidden, gin.H{"error": "You can only update your tasks"})
        return
    }

    updatedTask.ID = id
    updatedTask.UserID = task.UserID

    task, err = data.UpdateTask(id,updatedTask)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context){
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
        return
    }

	user,_ := c.Get("user")
	claims := user.(*middleware.Claims)
	
	task, err := data.GetTaskByID(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if task.UserID != claims.UserID && claims.Role != "admin"{
        c.JSON(http.StatusForbidden, gin.H{"error": "You can only delete your tasks"})
        return
    }

	if err := data.DeleteTask(id); err != nil{
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusNoContent, nil)
}