package controllers

import (
	"net/http"
	// "strconv"
	"task_manager/data"
	"task_manager/models"

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
	
	task,err := data.CreateTask(newtask)
	if err != nil{
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, task)
}

// func UpdateTask(c *gin.Context){
// 	id, err := primitive.ObjectIDFromHex(c.Param("id"))
// 	if err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
//         return
//     }

// 	var updatedTask models.Task
// 	if err := c.ShouldBindJSON(&updatedTask); err != nil{
// 		c.JSON(http.StatusNotFound, err.Error())
// 		return
// 	}

// 	task, err := data.UpdateTask(id, updatedTask)
// 	if err != nil{
// 		if err.Error() == "Task not found"{
// 			c.JSON(http.StatusNotFound, err.Error())
// 		}else{
// 			c.JSON(http.StatusInternalServerError, err.Error())
// 		}
// 		return	
// 	}
// 	c.JSON(http.StatusOK, task)
// }

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

	task, err := data.UpdateTask(id, updatedTask)
	if err != nil {
		if err.Error() == "No task found with the given ID" {
			c.JSON(http.StatusNotFound, err.Error())
		} else {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
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

	if err := data.DeleteTask(id); err != nil{
		c.JSON(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusNoContent, nil)
}