package main

import (
	"user-management-api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	route.GET("/", controllers.HomeController)

	route.Run()
}
