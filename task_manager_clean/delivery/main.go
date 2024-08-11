package main

import (
	"context"
	"log"
	"task_manager/config"
	"task_manager/repositories"
	"task_manager/usecases"
	"task_manager/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	client, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	userCollection := client.Database("TaskManager").Collection("Users")
	taskCollection := client.Database("TaskManager").Collection("Tasks")

	userRepo := repositories.NewUserRepository(userCollection)
	taskRepo := repositories.NewTaskRepository(taskCollection)

	userUsecase := usecase.NewUserUsecase(userRepo)
	taskUsecase := usecase.NewTaskUsecase(taskRepo)

	userController := controllers.NewUserController(userUsecase)
	taskController := controllers.NewTaskController(taskUsecase)

	r := gin.Default()

	publicRoutes := r.Group("/api")
	routers.SetPublicRoutes(publicRoutes, userController)

	protectedRoutes := r.Group("/api")
	routers.SetProtectedRoutes(protectedRoutes, taskController, userController)

	r.Run()
}
