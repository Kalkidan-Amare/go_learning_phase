package main

import (
	"log"
	"task_manager/router"
	"task_manager/data"
	"task_manager/config"

	"github.com/gin-gonic/gin"

)

func main(){
	client, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}
	
	data.SetUserCollection(client)
	data.SetTaskCollection(client)
	r := gin.Default()
	router.SetupTaskRouter(r)
	router.SetupUserRouter(r)

	r.Run()	
}