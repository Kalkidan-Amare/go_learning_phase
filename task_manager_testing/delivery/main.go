package main

import (
	"context"
	"log"
	"task_manager/config"
	"task_manager/delivery/controllers"
	"task_manager/delivery/routers"
	"task_manager/repositories"
	"task_manager/usecase"

	// "task_manager/controllers"

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

	userMockCollection := repositories.NewMongoCollection(userCollection)
	taskMockCollection := repositories.NewMongoCollection(taskCollection)

	userRepo := repositories.NewUserRepository(userMockCollection)
	taskRepo := repositories.NewTaskRepository(taskMockCollection)

	userUsecase := usecase.NewUserUsecase(userRepo)
	taskUsecase := usecase.NewTaskUsecase(taskRepo)

	userController := controllers.NewUserController(userUsecase)
	taskController := controllers.NewTaskController(taskUsecase)

	r := gin.Default()

	publicRoutes := r.Group("/")
	routers.SetPublicRoutes(publicRoutes, userController)

	protectedRoutes := r.Group("/tasks")
	routers.SetProtectedRoutes(protectedRoutes, taskController, userController)

	r.Run()
}
