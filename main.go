package main

import (
	"user-management-api/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	
	route := gin.Default()

	route.GET("/", controllers.Home)
	route.GET("/user/:UserID", controllers.GetUser)
	route.POST("/user", controllers.AddUser)
	route.GET("/auth/user/:UserID/password/:Password", controllers.AuthenticateUser)

	route.Run()
}