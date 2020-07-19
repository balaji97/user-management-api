package main

import (
	"log"
	"user-management-api/repository"
	"user-management-api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	initModules()
	
	route := gin.Default()

	route.GET("/", controllers.Home)
	route.GET("/user/:UserID", controllers.GetUser)
	route.POST("/user", controllers.AddUser)

	route.Run()
}

func initModules() {
	err := repository.InitializeRepository()
	if(err != nil) {
		log.Fatal(err)
	}
	controllers.InitializeControllers()
}