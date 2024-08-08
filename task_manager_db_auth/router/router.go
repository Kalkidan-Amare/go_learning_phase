package router

import (
    "task_manager/controllers"
    "task_manager/middleware"

	"github.com/gin-gonic/gin"
)

func SetupTaskRouter(r *gin.Engine) {
    authorized := r.Group("/")
    authorized.Use(middleware.AuthMiddleware())
    {
        authorized.GET("/tasks", controllers.GetTasks)
        authorized.GET("/tasks/:id", controllers.GetTask)
        authorized.POST("/tasks", controllers.CreatTask)

        authorized.PUT("/tasks/:id",middleware.AdminMiddleware(), controllers.UpdateTask)
        authorized.DELETE("/tasks/:id", middleware.AdminMiddleware(), controllers.DeleteTask)
    }
}

func SetupUserRouter(r *gin.Engine) {
    r.POST("/register", controllers.RegisterUser)
    r.POST("/register-admin", controllers.RegisterAdmin)
    r.POST("/login", controllers.LoginUser)
}