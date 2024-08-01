package controllers

import (
	"net/http"
	"task_manager/data"
	"task_manager/models"
	"strconv"
	"github.com/gin-gonic/gin"
)


func GetTasks(c *gin.Context){
	all := data.GetAllTasks()
	c.JSON(200, all)
}

func GetTask(c *gin.Context){
	id,err := strconv.Atoi(c.Param("id"))
	if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
        return
    }

	task,err := data.GetTaskByID(id)
	if err != nil{
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, task)
}

func CreatTask(c *gin.Context){
	var newtask models.Task
	if err := c.ShouldBindJSON(&newtask); err != nil{
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	
	c.JSON(http.StatusOK, data.CreateTask(newtask))
}

func UpdateTask(c *gin.Context){
	id,err := strconv.Atoi(c.Param("id"))
	if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
        return
    }

	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil{
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	task, err := data.UpdateTask(id, updatedTask)
	if err != nil{
		c.JSON(http.StatusBadRequest, err.Error())
		return	
	}
	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context){
	id,err := strconv.Atoi(c.Param("id"))
	if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
        return
    }

	if err := data.DeleteTask(id); err != nil{
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusNoContent, nil)
}