package controllers

import (
	"user-management-api/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Home - "/"
func Home(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"data": "Hello World!"})
}

// GetUser - "/user/{UserID}" - Return User as JSON object for given ID
func GetUser(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"data": repository.GetUser(context.Param("UserID"))})
}
