package routers

import (
	"task_manager/delivery/controllers"
	"task_manager/infrastructure"

	"github.com/gin-gonic/gin"
)

func SetPublicRoutes(r *gin.RouterGroup, userController *controllers.UserController) {
	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)
}

func SetProtectedRoutes(r *gin.RouterGroup, taskController *controllers.TaskController, userController *controllers.UserController) {
	r.Use(infrastructure.AuthMiddleware())
	r.GET("/", taskController.GetTasks)
	r.POST("/", taskController.CreateTask)
	r.GET("/:id", taskController.GetTask)
	r.PUT("/:id", taskController.UpdateTask)
	r.DELETE("/:id", taskController.DeleteTask)

	// r.GET("/users/:id", userController.GetUserByID)
	// r.GET("/users", userController.GetAllUsers)
	// r.PUT("/users/:id", userController.UpdateUser)
}
