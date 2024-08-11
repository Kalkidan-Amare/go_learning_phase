package routers

import (
	"task_manager/controllers"
	"task_manager/infrastructure"

	"github.com/gin-gonic/gin"
)

func SetPublicRoutes(r *gin.RouterGroup, userController *controllers.UserController) {
	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)
}

func SetProtectedRoutes(r *gin.RouterGroup, taskController *controllers.TaskController, userController *controllers.UserController) {
	r.Use(infrastructure.AuthMiddleware())
	r.GET("/tasks", taskController.GetTasks)
	r.POST("/tasks", taskController.CreateTask)
	r.GET("/tasks/:id", taskController.GetTask)
	r.PUT("/tasks/:id", taskController.UpdateTask)
	r.DELETE("/tasks/:id", taskController.DeleteTask)

	r.GET("/users", userController.GetAllUsers)
	r.GET("/users/:id", userController.GetUser)
	r.PUT("/users/:id", userController.UpdateUser)
}
